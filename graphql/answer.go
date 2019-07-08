package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/graphql-go/graphql"
	"github.com/piex/interview/protorepo/answer"
)

// AnswerListPageAble 分页
type AnswerListPageAble struct {
	TotalCount int64            `json:"totalCount"`
	Nodes      []*answer.Answer `json:"nodes"`
}

var ansFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "answer",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "answer ID",
			Type:        graphql.Int,
		},
		"quId": &graphql.Field{
			Description: "question ID",
			Type:        graphql.Int,
		},
		"content": &graphql.Field{
			Description: "content",
			Type:        graphql.String,
		},
		"createdAt": &graphql.Field{
			Description: "created at",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ans := p.Source.(*answer.Answer)
				return ans.CreatedAt.Seconds * 1000, nil
			},
		},
		"updatedAt": &graphql.Field{
			Description: "updated at",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ans := p.Source.(*answer.Answer)
				return ans.UpdatedAt.Seconds * 1000, nil
			},
		},
	},
})

var ansListFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ansList",
	Description: "",
	Fields: graphql.Fields{
		"totalCount": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "count of answers.",
		},
		"nodes": &graphql.Field{
			Description: "answer of the list.",
			Type:        graphql.NewList(ansFieldType),
		},
	},
})

var queryAnswerField = graphql.Field{
	Description: "get answer detail",
	Type:        ansFieldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "question id",
			Type:        graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := int64(p.Args["id"].(int))

		req := answer.GetAnswerRequest{
			Id: id,
		}

		ans, err := ansClient.GetAnswer(context.Background(), &req)

		if err != nil {
			glog.Error(err)
		}

		return ans, err
	},
}

var queryAnswerListField = graphql.Field{
	Description: "",
	Type:        ansListFieldType,
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Description: "页码",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"size": &graphql.ArgumentConfig{
			Description: "页数",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"quId": &graphql.ArgumentConfig{
			Description: "question id",
			Type:        graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		page := int32(p.Args["page"].(int))
		size := int32(p.Args["size"].(int))
		quID := int64(p.Args["quId"].(int))

		req := answer.GetAnswerListRequest{
			Page: page,
			Size: size,
			QuId: quID,
		}

		ansList, err := ansClient.GetAnswerList(context.Background(), &req)

		if err != nil {
			glog.Error(err)
		}

		res := AnswerListPageAble{
			TotalCount: ansList.Total,
			Nodes:      ansList.Answers,
		}

		return res, err
	},
}

var mutationCreateAnswer = graphql.Field{
	Description: "create a new answer",
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name:        "newAnswer",
		Description: "",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Description: "answer ID",
				Type:        graphql.Int,
			},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"quId": &graphql.ArgumentConfig{
			Description: "The id of question answer belongs to",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"content": &graphql.ArgumentConfig{
			Description: "The content of answer.",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		content := p.Args["content"].(string)
		quID := int64(p.Args["quId"].(int))

		ans := answer.Answer{
			QuId:    quID,
			Content: content,
		}

		res, err := ansClient.CreateAnswer(context.Background(), &ans)
		if err != nil {
			glog.Error(err)
		}

		return res, err
	},
}

var mutationUpdateAnswer = graphql.Field{
	Description: "update answer",
	Type:        successFieldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "The id of answer",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"content": &graphql.ArgumentConfig{
			Description: "The content of answer.",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ID := int64(p.Args["id"].(int))
		content := p.Args["content"].(string)

		ans := answer.Answer{
			Id:      ID,
			Content: content,
		}

		_, err := ansClient.UpdateAnswer(context.Background(), &ans)

		res := Success{
			Success: true,
		}

		if err != nil {
			glog.Error(err)
			res.Success = false
		}
		return res, err
	},
}
