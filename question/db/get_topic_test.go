package db

import "testing"

func TestGetTopic(t *testing.T) {
	topics, err := GetTopic(1)

	if err != nil {
		t.Error(err)
	}

	if topics.Total == 0 {
		t.Error("total is zero.")
	}

	t.Errorf("%+v\n", topics)
}
