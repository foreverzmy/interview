package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/question"
	"github.com/piex/interview/question/db"
)

// QuestionService grpc service
type QuestionService struct{}

// CreateQuestion 创建问题
func (s *QuestionService) CreateQuestion(ctx context.Context, qs *question.Question) (*question.CreateQuestionResponse, error) {

	id, err := db.CreateQuestion(qs)

	res := question.CreateQuestionResponse{
		Id: id,
	}

	return &res, err
}

// UpdateQuestion 更新问题
func (s *QuestionService) UpdateQuestion(ctx context.Context, qs *question.Question) (res *question.UpdateQuestionResponse, err error) {
	err = db.UpdateQuestion(qs)

	res.Success = true

	return
}

// GetQuestionDetail 查看详情
func (s *QuestionService) GetQuestionDetail(ctx context.Context, req *question.Question) (res *question.Question, err error) {
	qsOrm, err := db.GetQuestion(req.Id)

	if err != nil {
		return
	}

	qs, err := qsOrm.ToPB(ctx)

	return &qs, nil
}

// GetQuestionList 查看列表
func (s *QuestionService) GetQuestionList(ctx context.Context, req *question.GetQuestionListRequest) (*question.QuestionList, error) {

	qusOrm, total, err := db.GetQuestionList(req.Page, req.Size, req.Keyword)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	res := &question.QuestionList{
		Total:     total,
		Questions: nil,
	}

	for _, quOrm := range qusOrm {
		qu, err := quOrm.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		res.Questions = append(res.Questions, &qu)
	}

	return res, err
}
