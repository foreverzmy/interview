package db

import (
	"testing"

	"github.com/piex/interview/protorepo/question"
)

func TestUpdateQuestion(t *testing.T) {
	qs := question.Question{
		Id:         1,
		Title:      "语义化标签",
		Summary:    "什么是语义化标签，有什么作用，有哪些语义化标签？",
		Content:    "什么是语义化标签，有什么作用，有哪些语义化标签？",
		Difficulty: 1,
	}

	err := UpdateQuestion(&qs)

	if err != nil {
		t.Error(err)
	}
}
