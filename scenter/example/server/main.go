package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/foreverzmy/scenter"
	"github.com/foreverzmy/scenter/example/api"
	"google.golang.org/grpc"
)

var (
	serv = flag.String("service", "ping_service", "service name")
	host = flag.String("host", "localhost", "listening host")
	// port = flag.String("port", "50001", "listening port")
	reg = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	port := flag.String("port", "", "Input service port")

	flag.Parse()

	lis, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		panic(err)
	}

	err = scenter.Register(*serv, *host, *port, *reg, time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		scenter.UnRegister()
		os.Exit(1)
	}()

	log.Printf("starting hello service at %s", *port)

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &PingServer{})
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		scenter.UnRegister()
	}
}

// PingServer represents the gRPC server
type PingServer struct{}

// SayHello generates response to a Ping request
func (s *PingServer) SayHello(ctx context.Context, in *api.PingMessage) (*api.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &api.PingMessage{Greeting: "bar"}, nil
}
