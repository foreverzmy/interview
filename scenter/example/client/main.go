package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/foreverzmy/scenter"
	"github.com/foreverzmy/scenter/example/api"
	"google.golang.org/grpc"
)

var (
	serv = flag.String("service", "ping_service", "service name")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()
	r := scenter.NewResolver(*serv)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()

	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	for t := range ticker.C {
		client := api.NewPingClient(conn)
		resp, err := client.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
		if err != nil {
			fmt.Printf("%v: Reply is %s\n", t, err)
		}

		log.Printf("Response from server: %s", resp.Greeting)
	}
}
