package db

import (
	"github.com/piex/interview/protorepo/question"
)

// GetQuestionList 问题列表
func GetQuestionList(page int32, size int32, keyword string) (questions []question.QuestionORM, total int64, err error) {
	err = DB.Table("question").
		Where("title LIKE ?", "%"+keyword+"%").
		Count(&total).
		Error

	if err != nil {
		return
	}

	err = DB.Table("question").
		Where("title LIKE ?", "%"+keyword+"%").
		Limit(size).Offset((page - 1) * size).
		Find(&questions).
		Error

	return
}
