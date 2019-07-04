package db

import (
	"testing"

	"github.com/piex/interview/protorepo/question"
)

func TestUpdateQuestion(t *testing.T) {
	qs := question.Question{
		Id:         1,
		Title:      "什么是闭包？",
		Summary:    "什么是闭包？",
		Content:    "什么是闭包？",
		Difficulty: 0,
	}

	err := UpdateQuestion(&qs)

	if err != nil {
		t.Error(err)
	}
}
