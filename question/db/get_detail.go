package db

import (
	"github.com/piex/interview/protorepo/golang/question"
)

// GetQuestion 查询问题详情
func GetQuestion(id int64) (qs question.QuestionORM, err error) {

	err = DB.Table("question").Where("id = ?", id).Find(&qs).Error

	return
}
