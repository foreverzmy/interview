package db

import "testing"

func TestGetTopic(t *testing.T) {
	topics, err := GetTopic(1)

	if err != nil {
		t.Error(err)
	}

	t.Errorf("%+v\n", topics)
}
