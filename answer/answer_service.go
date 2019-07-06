package main

import (
	"context"

	"github.com/piex/interview/answer/db"
	"github.com/piex/interview/protorepo/answer"
)

// AnswerService Answer Service
type AnswerService struct {
}

// CreateAnswer 创建
func (s *AnswerService) CreateAnswer(ctx context.Context, ans *answer.Answer) (*answer.CreateAnswerResponse, error) {
	id, err := db.CreateAnswer(ans)

	res := answer.CreateAnswerResponse{
		Id: id,
	}

	return &res, err
}

// UpdateAnswer 编辑
func (s *AnswerService) UpdateAnswer(ctx context.Context, ans *answer.Answer) (*answer.Empty, error) {
	ansORM, err := ans.ToORM(ctx)

	err = db.UpdateAnswer(&ansORM)

	res := answer.Empty{}

	return &res, err
}

// GetAnswerList 获取列表
func (s *AnswerService) GetAnswerList(ctx context.Context, req *answer.GetAnswerListRequest) (*answer.AnswerList, error) {
	ansORMList, total, err := db.GetAnswerList(req.Page, req.Size, req.QuId)

	res := answer.AnswerList{
		Total:   total,
		Answers: nil,
	}

	res.Total = total

	for _, ansORM := range ansORMList {
		ans, err := ansORM.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		res.Answers = append(res.Answers, &ans)
	}

	return &res, err
}
