package db

import (
	"context"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/question"
)

// GetQusByTopic 根据 topic 查询 问题列表
func GetQusByTopic(page int32, size int32, topicIDs int64) (quList question.QuestionList, err error) {

	var qusORM []*question.QuestionORM

	err = DB.Table("qu_topic").Where("topic_id = ?", topicIDs).Count(&quList.Total).Error

	if err != nil {
		return
	}

	err = DB.Table("question").Joins("left join qu_topic on qu_topic.topic_id = question.id").Where("qu_topic.topic_id = ?", topicIDs).Limit(size).Offset((page - 1) * size).Find(&qusORM).Error

	for _, quORM := range qusORM {
		var pb question.Question
		pb, err = quORM.ToPB(context.Background())
		quList.Questions = append(quList.Questions, &pb)

		if err != nil {
			glog.Error(pb)
			return
		}
	}

	return
}
