package array

// IsAnagram https://leetcode.com/problems/valid-anagram/description/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sChecker := make(map[rune]int)
	tChecker := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		sChecker[rune(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		tChecker[rune(t[i])]++
	}
	for r, v := range sChecker {
		if tChecker[r] != v {
			return false
		}
	}
	return true
}
