package db

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/question"
)

// CreateQuestion 创建问题
func CreateQuestion(qs *question.Question) (id int64, err error) {
	orm, err := qs.ToORM(nil)

	err = DB.Table("question").Create(&orm).Error

	if err != nil {
		glog.Error(err)
		fmt.Print(err)
	}

	return orm.Id, err
}
