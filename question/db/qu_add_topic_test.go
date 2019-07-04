package db

import "testing"

func TestQuAddTopic(t *testing.T) {
	topics := []int64{2}

	err := QuAddTopic(2, topics)

	if err != nil {
		t.Error(err)
	}
}
