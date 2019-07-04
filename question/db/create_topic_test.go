package db

import (
	"testing"

	"github.com/piex/interview/protorepo/topic"
)

func TestCreateTopic(t *testing.T) {
	topic := topic.Topic{
		Slug: "对象",
	}

	id, err := CreateTopic(&topic)

	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}
