package db

import (
	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/question"
)

// UpdateQuestion 更新问题
func UpdateQuestion(qu *question.QuestionORM) error {

	err := DB.Table("question").Save(qu).Error

	if err != nil {
		glog.Error(err)
	}

	return err
}
