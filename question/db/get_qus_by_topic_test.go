package db

import "testing"

func TestGetQusByTopic(t *testing.T) {
	quList, err := GetQusByTopic(1, 10, 1)

	if err != nil {
		t.Error(err)
	}

	t.Errorf("%+v\n", quList)
}
