package db

import (
	"testing"

	"github.com/piex/interview/protorepo/answer"
)

func TestCreateAnswer(t *testing.T) {
	ans := answer.Answer{
		QuId:    1,
		Content: "200 ok",
	}

	id, err := CreateAnswer(&ans)

	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}
