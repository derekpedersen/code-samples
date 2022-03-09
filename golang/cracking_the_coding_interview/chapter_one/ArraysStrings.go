package chapter_one

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
