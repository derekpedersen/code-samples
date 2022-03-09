package chapter_one_test

import (
	"testing"

	"github.com/derekpedersen/code-samples/golang/cracking_the_coding_interview/chapter_one"
)

// 1.1
// implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
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

// 1.2
// implement a function void reverse(char* str) in C which reverses a null-terminated string
func Test_ChapterOne_Two(t *testing.T) {
	testTable := []struct {
		Val    string
		Result string
	}{
		{
			Val:    "asdf",
			Result: "fdsa",
		},
		{
			Val:    "derek",
			Result: "kered",
		},
	}

	for _, v := range testTable {
		result := chapter_one.ReverseString(v.Val)
		if result != v.Result {
			t.Errorf("expected %v but got %v for string %v", v.Result, result, v.Val)
		}
	}
}

// 1.3
// given two strings write a method to decide if one is a permutation of the other
func Test_ChapterOne_Three(t *testing.T) {
	testTable := []struct {
		One    string
		Two    string
		Result bool
	}{
		{
			One:    "asdf",
			Two:    "fdsa",
			Result: true,
		},
		{
			One:    "derek",
			Two:    "keed",
			Result: false,
		},
		{
			One:    "fast",
			Two:    "last",
			Result: false,
		},
	}

	for _, v := range testTable {
		result := chapter_one.ArePermutations(v.One, v.Two)
		if result != v.Result {
			t.Errorf("expected %v but got %v for strings %v and %v", v.Result, result, v.One, v.Two)
		}
	}
}
