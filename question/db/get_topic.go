package db

import (
	"context"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/topic"
)

// GetTopic 查找 qu 的 topic
func GetTopic(quID int64) (topics topic.TopicList, err error) {

	var tocs []topic.TopicORM

	tx := DB.Begin()

	err = tx.Table("topic").Joins("left join qu_topic on topic.id = qu_topic.topic_id").Where("qu_topic.qu_id = ?", quID).Count(&topics.Total).Error

	err = tx.Table("topic").Joins("left join qu_topic on topic.id = qu_topic.topic_id").Where("qu_topic.qu_id = ?", quID).Find(&tocs).Error

	tx.Commit()

	for _, t := range tocs {
		var pb topic.Topic
		pb, err = t.ToPB(context.Background())
		topics.Topics = append(topics.Topics, &pb)

		if err != nil {
			glog.Error(pb)
			return
		}
	}

	return
}
