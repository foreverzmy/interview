package db

import "testing"

func TestGetTopicList(t *testing.T) {
	topics, total, err := GetTopicList()

	if err != nil {
		t.Error(err)
	}

	if total == 0 {
		t.Error("topics number is zero.")
	}

	t.Log((topics))
}
