package db

import (
	"testing"

	"github.com/piex/interview/protorepo/answer"
)

func TestUpdateAnswer(t *testing.T) {
	ans := answer.AnswerORM{
		Id:      1,
		Content: "ssd",
	}

	err := UpdateAnswer(&ans)

	if err != nil {
		t.Error(err)
	}

}
