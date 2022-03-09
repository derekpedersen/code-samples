package chapter_one

import "sort"

// 1.1
// implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
func IsUnique(val string) bool {
	used := make(map[rune]int)
	for _, v := range val {
		if _, ok := used[v]; ok {
			return false
		}
		used[v] = 1
	}

	return true
}

// 1.2
// implement a function void reverse(char* str) in C which reverses a null-terminated string
func ReverseString(val string) string {
	chars := []rune(val)
	responsechars := []rune{}

	for i := len(chars) - 1; i >= 0; i-- {
		responsechars = append(responsechars, chars[i])
	}

	return string(responsechars)
}

// 1.3
// given two strings write a method to decide if one is a permutation of the other
func ArePermutations(one string, two string) bool {
	ones := []rune(one)
	twos := []rune(two)

	if len(ones) != len(twos) {
		return false
	}

	sort.Slice(ones, func(i int, j int) bool {
		return ones[i] < ones[j]
	})

	sort.Slice(twos, func(i int, j int) bool {
		return twos[i] < twos[j]
	})

	for i := 0; i < len(ones); i++ {
		if ones[i] != twos[i] {
			return false
		}
	}

	return true
}
