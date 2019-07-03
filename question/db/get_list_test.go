package db

import "testing"

func TestGetQuestionList(t *testing.T) {
	questions, total, err := GetQuestionList(1, 10, "")

	if err != nil {
		t.Error(err)
	}

	if total == 0 {
		t.Errorf("error with total = 0: %+v", questions)
	}

}
