package db

import "testing"

func TestGetQuestion(t *testing.T) {
	qs, err := GetQuestion(1)

	if err != nil {
		t.Error(err)
	}

	if qs.Id != 1 {
		t.Error(qs)
	}

	t.Log(qs)
}
