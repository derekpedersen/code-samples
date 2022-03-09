package chapter_one

import (
	"sort"
	"strconv"
)

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

// 1.4
// write a method to replace all spaces with `%20`
func EncodeSpaces(val string) string {
	encoded := []rune(`%20`)
	result := []rune{}
	for _, v := range val {
		if v == ' ' {
			result = append(result, encoded...)
		} else {
			result = append(result, v)
		}
	}

	return string(result)
}

// 1.5
// implement a method to perform a basic string compression using the counts of repeated characters
func CompressString(val string) string {
	counts := make(map[rune]int)
	vals := []rune(val)

	sort.Slice(vals, func(i int, j int) bool {
		return vals[i] < vals[j]
	})

	for k, v := range vals {
		if _, ok := counts[v]; ok {
			continue
		}
		// check if next one is the same
		for i := 1; i > 0; i++ {
			if len(vals) > k+i && v == vals[k+i] {
				continue
			} else {
				counts[v] = i
				i = -1
			}
		}
	}

	str := []rune{}
	for k, v := range counts {
		count := []rune(string(k) + strconv.FormatInt(int64(v), 10))
		str = append(str, count...)
	}

	response := string(str)

	if len(response) > len(val) {
		return val
	}

	return response
}
