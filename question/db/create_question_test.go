package db

import (
	"testing"

	"github.com/piex/interview/protorepo/question"
)

func TestCreateQuestion(t *testing.T) {
	qs := question.Question{
		Title:      "什么是闭包？",
		Summary:    "什么是闭包？",
		Content:    "什么是闭包？",
		Difficulty: 2,
	}

	id, err := CreateQuestion(&qs)

	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}
