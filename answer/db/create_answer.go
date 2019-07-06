package db

import (
	"context"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/answer"
)

// CreateAnswer 创建 answer
func CreateAnswer(ans *answer.Answer) (int64, error) {
	orm, err := ans.ToORM(context.Background())

	err = DB.Table("answer").Create(&orm).Error

	if err != nil {
		glog.Error(err)
	}

	return orm.Id, err
}
