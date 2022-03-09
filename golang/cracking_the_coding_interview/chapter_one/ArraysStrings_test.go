package chapter_one_test

import (
	"testing"

	"github.com/derekpedersen/code-samples/golang/cracking_the_coding_interview/chapter_one"
)

func Test_ChapterOne_One(t *testing.T) {
	testTable := []struct {
		Val    string
		Unique bool
	}{
		{
			Val:    "asdf",
			Unique: true,
		},
		{
			Val:    "derek",
			Unique: false,
		},
	}

	for _, v := range testTable {
		result := chapter_one.IsUnique(v.Val)
		if result != v.Unique {
			t.Errorf("expected %v but got %v for string %v", v.Unique, result, v.Val)
		}
	}
}
