package main

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"question":  &queryQuestionField,
		"questions": &queryQuestionListField,
		"topics":    &queryTopicListField,
		"topic":     &queryTopicField,
		"answers":   &queryAnswerListField,
	},
})
