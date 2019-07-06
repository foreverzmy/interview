package db

import (
	"github.com/piex/interview/protorepo/topic"
)

// GetTopic 根据 id 查找 topic
func GetTopic(id int64) (toc topic.TopicORM, err error) {

	err = DB.Table("topic").Where("id = ?", id).Find(&toc).Error

	return
}
