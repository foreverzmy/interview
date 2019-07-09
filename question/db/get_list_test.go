package db

import "testing"

func TestGetQuestionList(t *testing.T) {
	questions, total, err := GetQuestionList(1, 10, "特点与判断")

	if err != nil {
		t.Error(err)
	}

	if total == 0 {
		t.Errorf("error with total = 0: %+v", questions)
	}

	t.Error(questions)

}
