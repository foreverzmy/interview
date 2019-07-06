package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/graphql-go/graphql"
	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/protorepo/topic"
)

// TopicList topic list
type TopicList struct {
	TotalCount int64          `json:"totalCount"`
	Nodes      []*topic.Topic `json:"nodes"`
}

var topicFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "topic",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "question ID",
			Type:        graphql.Int,
		},
		"slug": &graphql.Field{
			Description: "feed title",
			Type:        graphql.String,
		},
	},
})

var topicListFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "topics",
	Description: "list of topic",
	Fields: graphql.Fields{
		"totalCount": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Count of all topics.",
		},
		"nodes": &graphql.Field{
			Description: "nodes of the list.",
			Type:        graphql.NewList(topicFieldType),
		}},
})

var queryTopicField = graphql.Field{
	Description: "topic",
	Type:        topicListFieldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "The id topic",
			Type:        graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

	},
}

var queryTopicListField = graphql.Field{
	Description: "A list of topic",
	Type:        topicListFieldType,
	Args: graphql.FieldConfigArgument{
		"quId": &graphql.ArgumentConfig{
			Description: "The id of question topics belongs to",
			Type:        graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		quIDArg, ok := p.Args["quId"].(int)

		var quID int64

		if ok == true {
			quID = int64(quIDArg)
		} else {
			qu, ok := p.Source.(*question.Question)

			if ok == true {
				quID = qu.Id
			}
		}

		req := topic.GetTopicRequest{
			QuId: quID,
		}

		var topics *topic.TopicList
		var err error

		if quID > 0 {
			topics, err = topicClient.GetTopic(context.Background(), &req)
		} else {
			topics, err = topicClient.GetTopicList(context.Background(), &topic.Empty{})
		}

		res := TopicList{
			TotalCount: topics.Total,
			Nodes:      topics.Topics,
		}

		return res, err
	},
}

// 创建 topic
var mutaionCreateTopic = graphql.Field{
	Description: "Create a new topic",
	Type:        topicFieldType,
	Args: graphql.FieldConfigArgument{
		"slug": &graphql.ArgumentConfig{
			Description: "The slug of topic.",
			Type:        graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		slug := p.Args["slug"].(string)

		topic := topic.Topic{
			Slug: slug,
		}

		res, err := topicClient.CreateTopic(context.Background(), &topic)

		if err != nil {
			glog.Error(err)
		}

		return res, err
	},
}