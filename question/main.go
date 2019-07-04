package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/scenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var lis net.Listener
var err error

// glog 初始化
func initGlog() {
	//  初始化 glog，主要使服务器启动后自己直接加载，并不用命令行执行对应的参数
	flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", "./log")        // 日志文件保存目录
	flag.Set("v", "5")                  // 配置V输出的等级。
	flag.Parse()
}

func initGrpc() {
	flag.Parse()
	serv := flag.String("service", "question_service", "service name")
	host := flag.String("host", "localhost", "listening host")
	port := flag.String("port", "50001", "listening port")
	reg := flag.String("reg", "http://localhost:2379", "register etcd address")
	lis, err = net.Listen("tcp", net.JoinHostPort(*host, *port))
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

	log.Printf("starting service at %s", *port)
}

func init() {
	initGlog()
	initGrpc()
}

func main() {
	// create a gRPC server object
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}))

	question.RegisterQuestionServiveServer(grpcServer, &QuestionService{})

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		scenter.UnRegister()
	}
}
