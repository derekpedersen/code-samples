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

// 1.4
// write a method to replace all spaces with `%20`
func Test_ChapterOne_Four(t *testing.T) {
	testTable := []struct {
		Val    string
		Result string
	}{
		{
			Val:    "as df",
			Result: "as%20df",
		},
		{
			Val:    "derek the great",
			Result: "derek%20the%20great",
		},
	}

	for _, v := range testTable {
		result := chapter_one.EncodeSpaces(v.Val)
		if result != v.Result {
			t.Errorf("expected %v but got %v for string %v", v.Result, result, v.Val)
		}
	}
}

// 1.5
// implement a method to perform a basic string compression using the counts of repeated characters
func Test_ChapterOne_Five(t *testing.T) {
	testTable := []struct {
		Val    string
		Result string
	}{
		{
			Val:    "aaaa",
			Result: "a4",
		},
		{
			Val:    "aaabbccccdd",
			Result: "a3b2c4d2",
		},
		{
			Val:    "a",
			Result: "a", // as a1 is longer than just a so no need to compress
		},
	}

	for _, v := range testTable {
		result := chapter_one.CompressString(v.Val)
		if result != v.Result {
			t.Errorf("expected %v but got %v for string %v", v.Result, result, v.Val)
		}
	}
}

// 1.6
// given an image represented by an NxN matrix where each pixel in teh image is 4 bytes write a method to rotate the image by 90 degrees

// 1.7
// write an algorithm such taht if an element in an NxN matrix is 0 its entire row and column are set to 0

// 1.8
// assume you have a method isSubstring which checks if one worse is a substring of another. given s1 and s2 write code to check if s2 is a rotation of s1 using only one call to isSubstring
