package db

import "testing"

func TestQuAddTopic(t *testing.T) {
	topics := []int64{1}

	err := QuAddTopic(1, topics)

	if err != nil {
		t.Error(err)
	}
}
