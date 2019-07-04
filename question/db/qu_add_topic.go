package db

import (
	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/topic"
)

// QuAddTopic qu 添加 topic
func QuAddTopic(qu int64, topics []int64) error {
	tx := DB.Begin()

	err := tx.Table("qu_topic").Delete(nil, "qu_id = ?", qu).Error

	if err != nil {
		glog.Error(err)
		return err
	}

	for _, v := range topics {
		err = tx.Table("qu_topic").Create(&topic.QuTopicORM{
			QuId:    qu,
			TopicId: v,
		}).Error
	}

	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return err
}
