package db

import "testing"

func TestGetTopic(t *testing.T) {
	topic, err := GetTopic(1)

	if err != nil {
		t.Error(err)
	}

	t.Log(topic)
}
