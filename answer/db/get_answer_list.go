package db

import (
	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/answer"
)

// GetAnswerList 根据 question id 查询 answer list
func GetAnswerList(page int32, size int32, quID int64) (anssORM []answer.AnswerORM, total int64, err error) {

	tx := DB.Begin()

	err = tx.Table("answer").
		Where("qu_id = ?", quID).
		Count(&total).
		Error

	err = DB.Table("answer").
		Where("qu_id = ?", quID).
		Limit(size).Offset((page - 1) * size).
		Find(&anssORM).
		Error

	if err != nil {
		glog.Error(err)
		tx.Rollback()
	}

	tx.Commit()

	return
}
