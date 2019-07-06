package main

import "github.com/graphql-go/graphql"

var successFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "success",
	Description: "",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Description: "state",
			Type:        graphql.Boolean,
		},
	},
})

// Success 成果
type Success struct {
	success bool
}
