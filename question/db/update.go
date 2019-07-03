package db

import (
	"github.com/piex/interview/protorepo/golang/question"
)

// UpdateQuestion 更新问题
func UpdateQuestion(qs *question.Question) (err error) {
	orm, err := qs.ToORM(nil)

	err = DB.Table("question").Save(orm).Error

	return
}
