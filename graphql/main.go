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

	"github.com/piex/interview/protorepo/answer"
	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/protorepo/topic"
	"github.com/piex/interview/scenter"
)

// Schema 类型
var (
	Schema graphql.Schema
	h      *handler.Handler

	qsClient    question.QuestionServiveClient
	topicClient topic.TopicServiceClient
	ansClient   answer.AnswerServiceClient

	reg = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
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

func getConn(serv *string) *grpc.ClientConn {
	r := scenter.NewResolver(*serv)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()

	if err != nil {
		panic(err)
	}

	return conn

}

func init() {
	quServ := flag.String("question_service", "question_service", "service name")
	ansServ := flag.String("answer_service", "answer_service", "service name")

	initGraphql()
	quConn := getConn(quServ)
	ansConn := getConn(ansServ)

	qsClient = question.NewQuestionServiveClient(quConn)
	topicClient = topic.NewTopicServiceClient(quConn)

	ansClient = answer.NewAnswerServiceClient(ansConn)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("Server Start.")

	r := NewRouter()
	http.ListenAndServe(":8099", r)
}
