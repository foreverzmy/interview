package db

import "testing"

func TestGetAnswer(t *testing.T) {
	ans, err := GetAnswer(2)

	if err != nil {
		t.Error(err)
	}

	t.Log(ans)
}
