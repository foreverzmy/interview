package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/graphql-go/graphql"
	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/protorepo/topic"
)

// QuestionListPageAble qs list 分页
type QuestionListPageAble struct {
	TotalCount int64                `json:"totalCount"`
	Nodes      []*question.Question `json:"nodes"`
}

var questionFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "question",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "question ID",
			Type:        graphql.Int,
		},
		"title": &graphql.Field{
			Description: "title",
			Type:        graphql.String,
		},
		"summary": &graphql.Field{
			Description: "summary",
			Type:        graphql.String,
		},
		"content": &graphql.Field{
			Description: "content",
			Type:        graphql.String,
		},
		"difficulty": &graphql.Field{
			Description: "difficulty",
			Type:        graphql.Int,
		},
		"topics": &queryTopicListField,
	},
})

var questionListFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "questions",
	Description: "list of questions",
	Fields: graphql.Fields{
		"totalCount": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "count of all questions.",
		},
		"nodes": &graphql.Field{
			Description: "Items of the list.",
			Type:        graphql.NewList(questionFieldType),
		},
	},
})

var queryQuestionField = graphql.Field{
	Description: "Lookup a given feed by id.",
	Type:        questionFieldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "The id of the feed.",
			Type:        graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := int64(p.Args["id"].(int))

		res, err := qsClient.GetQuestionDetail(context.Background(), &question.Question{Id: id})

		if err != nil {
			glog.Error(err)
		}

		return res, err
	},
}

var queryQuestionListField = graphql.Field{
	Description: "A list of questions",
	Type:        questionListFieldType,
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Description: "页码",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"size": &graphql.ArgumentConfig{
			Description: "页数",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"keyword": &graphql.ArgumentConfig{
			Description:  "关键字",
			Type:         graphql.String,
			DefaultValue: "",
		},
		"topicId": &graphql.ArgumentConfig{
			Description: "topic id",
			Type:        graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		page := int32(p.Args["page"].(int))
		size := int32(p.Args["size"].(int))
		keyword := p.Args["keyword"].(string)
		topicIDInt, topicIDOK := p.Args["topicId"].(int)

		t, tOK := p.Source.(*topic.Topic)

		var topicID int64

		topicID = int64(topicIDInt)

		if tOK {
			topicID = t.Id
		}

		if topicIDOK || tOK {

			req := topic.GetQusByTopicRequest{
				Page:    page,
				Size:    size,
				TopicId: topicID,
			}

			qus, err := topicClient.GetQusByTopic(context.Background(), &req)

			if err != nil {
				glog.Error((err))
			}

			res := QuestionListPageAble{
				TotalCount: qus.Total,
				Nodes:      qus.Questions,
			}

			return res, err
		}

		req := question.GetQuestionListRequest{
			Page:    page,
			Size:    size,
			Keyword: keyword,
		}

		qus, err := qsClient.GetQuestionList(context.Background(), &req)

		if err != nil {
			glog.Error((err))
		}

		res := QuestionListPageAble{
			TotalCount: qus.Total,
			Nodes:      qus.Questions,
		}

		return res, err
	},
}

var mutationCreateQuestion = graphql.Field{
	Description: "Create a new question",
	Type:        questionFieldType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Description: "The title of question.",
			Type:        graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Description: "The content of question.",
			Type:        graphql.String,
		},
		"difficulty": &graphql.ArgumentConfig{
			Description: "The difficulty of question.",
			Type:        graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		title := p.Args["title"].(string)
		content := p.Args["content"].(string)
		difficulty := int32(p.Args["difficulty"].(int))

		qs := question.Question{
			Title:      title,
			Content:    content,
			Difficulty: difficulty,
		}

		res, err := qsClient.CreateQuestion(context.Background(), &qs)

		if err != nil {
			glog.Error(err)
		}

		return res, err
	},
}
