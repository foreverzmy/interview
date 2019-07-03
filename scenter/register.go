package scenter

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	etcd3 "go.etcd.io/etcd/clientv3"
)

// 服务注册

// Prefix should start and end with no slash
var Prefix = "interview name"
var stopSignal = make(chan bool, 1)

// Deregister deregist service
var Deregister = make(chan struct{})

// Register register service
func Register(name, host, port string, target string, interval time.Duration, ttl int) error {
	serviceValue := net.JoinHostPort(host, port)
	serviceKey := fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceValue)
	// get endpoints for register dial address
	var err error
	client, err := etcd3.New(etcd3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		return fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
	}

	resp, err := client.Grant(context.TODO(), int64(ttl))
	if err != nil {
		return fmt.Errorf("grpclb: create etcd3 lease failed: %v", err)
	}

	if _, err := client.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
		return fmt.Errorf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
	}

	if _, err := client.KeepAlive(context.TODO(), resp.ID); err != nil {
		return fmt.Errorf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
	}

	// wait deregister then delete
	go func() {
		<-Deregister
		client.Delete(context.Background(), serviceKey)
		Deregister <- struct{}{}
	}()
	return nil
}

// UnRegister delete registered service from etcd
func UnRegister() {
	Deregister <- struct{}{}
	<-Deregister
}
