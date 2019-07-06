package db

import "testing"

func TestGetTopicByQu(t *testing.T) {
	topics, err := GetTopicByQu(1)

	if err != nil {
		t.Error(err)
	}

	if topics.Total == 0 {
		t.Error("total is zero.")
	}

	t.Errorf("%+v\n", topics)
}
