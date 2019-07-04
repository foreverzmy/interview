package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"google.golang.org/grpc"

	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/scenter"
)

// Schema 类型
var (
	Schema graphql.Schema
	h      *handler.Handler

	qsClient question.QuestionServiveClient

	serv = flag.String("service", "question_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func initGraphql() {
	var err error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)

	if err != nil {
		panic(err)
	}

	h = handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
}

func initGrpc() {
	r := scenter.NewResolver(*serv)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()

	if err != nil {
		panic(err)
	}

	qsClient = question.NewQuestionServiveClient(conn)
}

func init() {
	initGraphql()
	initGrpc()
}

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("Server Start.")

	r := NewRouter()
	http.ListenAndServe(":8099", r)
}
