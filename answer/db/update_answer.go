package db

import (
	"github.com/golang/glog"
	"github.com/piex/interview/protorepo/answer"
)

// UpdateAnswer 更新 answer
func UpdateAnswer(ans *answer.AnswerORM) error {
	var oldAns answer.AnswerORM

	err := DB.Table("answer").Where("id = ?", ans.Id).Find(&oldAns).Error

	ans.QuId = oldAns.QuId

	err = DB.Table("answer").Save(ans).Error

	if err != nil {
		glog.Error(err)
	}

	return err
}
