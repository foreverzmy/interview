package db

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/topic"
)

// CreateTopic 创建问题
func CreateTopic(t *topic.Topic) (id int64, err error) {
	orm, err := t.ToORM(nil)

	err = DB.Table("topic").Create(&orm).Error

	if err != nil {
		glog.Error(err)
		fmt.Print(err)
	}

	return orm.Id, err
}
