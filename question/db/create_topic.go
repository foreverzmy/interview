package db

import (
	"context"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/topic"
)

// CreateTopic 创建问题
func CreateTopic(t *topic.Topic) (int64, error) {
	orm, err := t.ToORM(context.Background())

	err = DB.Table("topic").Create(&orm).Error

	if err != nil {
		glog.Error(err)
	}

	return orm.Id, err
}
