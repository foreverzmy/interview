package db

import "testing"

func TestGetAnswerList(t *testing.T) {
	asnORMList, total, err := GetAnswerList(1, 10, 1)

	if err != nil {
		t.Error(err)
	}

	if total == 0 {
		t.Error("answers should not be zero")
	}

	t.Logf("%+v\n", asnORMList)
}
