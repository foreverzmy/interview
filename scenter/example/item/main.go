package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/foreverzmy/protorepo/golang/rsscomb"

	"github.com/foreverzmy/scenter"
	"google.golang.org/grpc"
)

var (
	serv = flag.String("service", "rsscomb_service", "service name")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func connectService(serviceName string) (*grpc.ClientConn, error) {
	r := scenter.NewResolver(serviceName)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	conn, _ := connectService(*serv)

	client := rsscomb.NewItemServiceClient(conn)
	resp, err := client.GetItemDetail(context.Background(), &rsscomb.GetItemDetailRequest{Id: 1})

	if err != nil {
		fmt.Printf("Reply is %s\n", err)
	}
	fmt.Printf("%+v\n", resp.Description)
}
