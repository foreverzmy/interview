package db

import (
	question "github.com/piex/interview/protorepo/topic"
)

// GetTopicList topic 列表
func GetTopicList() (topics []question.TopicORM, total int64, err error) {
	err = DB.Table("topic").Count(&total).Error

	if err != nil {
		return
	}

	err = DB.Table("topic").Find(&topics).Error

	return
}
