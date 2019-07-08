package main

import "github.com/graphql-go/graphql"

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createQuestion": &mutationCreateQuestion,
		"updateQuestion": &mutationUpdateQuestion,
		"createTopic":    &mutaionCreateTopic,
		"createAnswer":   &mutationCreateAnswer,
		"updateAnswer":   &mutationUpdateAnswer,
	},
})
