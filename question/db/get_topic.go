package db

import "github.com/piex/interview/protorepo/topic"

// GetTopic 查找 qu 的 topic
func GetTopic(quID int64) ([]topic.TopicORM, error) {
	var topics []topic.TopicORM

	err := DB.Table("topic").Joins("left join qu_topic on topic.id = qu_topic.topic_id").Where("qu_topic.qu_id = ?", quID).Find(&topics).Error

	return topics, err
}
