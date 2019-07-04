package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/graphql-go/graphql"
	"github.com/piex/interview/protorepo/topic"
)

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
