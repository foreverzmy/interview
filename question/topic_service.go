package main

import (
	"context"

	"github.com/piex/interview/protorepo/topic"
	"github.com/piex/interview/question/db"
)

// TopicService topic 服务
type TopicService struct {
}

// CreateTopic 创建 topic
func (s *TopicService) CreateTopic(ctx context.Context, req *topic.Topic) (*topic.CreateTopicResponse, error) {
	id, err := db.CreateTopic(req)

	res := topic.CreateTopicResponse{
		Id: id,
	}

	return &res, err
}

// GetTopicList 查找全部 topi
func (s *TopicService) GetTopicList(ctx context.Context, req *topic.Empty) (*topic.TopicList, error) {
	topics, total, err := db.GetTopicList()

	res := topic.TopicList{
		Total:  total,
		Topics: topics,
	}

	return &res, err
}

// GetTopic 查询 qu 的 topi
func (s *TopicService) GetTopic(ctx context.Context, req *topic.GetTopicRequest) (*topic.TopicList, error) {
	return nil, nil
}

// GetQusByTopic 查找有指定 topic 的 qus
func (s *TopicService) GetQusByTopic(ctx context.Context, req *topic.GetQusByTopicRequest) (*topic.QuestionList, error) {
	return nil, nil
}

// AddTopicsToQuestion qu 添加 topic
func (s *TopicService) AddTopicsToQuestion(ctx context.Context, req *topic.AddTopicsToQuestionRequest) (*topic.Empty, error) {
	return nil, nil
}
