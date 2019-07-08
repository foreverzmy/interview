package db

import "github.com/piex/interview/protorepo/answer"

// GetAnswer 查询回答详情
func GetAnswer(id int64) (ans answer.AnswerORM, err error) {
	err = DB.Table("answer").Where("id = ?", id).Find(&ans).Error

	return
}
