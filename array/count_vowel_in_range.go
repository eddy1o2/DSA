package array

// VowelStrings https://leetcode.com/problems/count-vowel-strings-in-ranges/?envType=daily-question&envId=2025-01-02
func VowelStrings(words []string, queries [][]int) []int {
	totalVowels := 0
	vowels := make([]int, len(words))
	for i, word := range words {
		if isVowel(word) {
			totalVowels++
		}
		vowels[i] = totalVowels
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		if query[0] == 0 {
			res[i] = vowels[query[1]]
			continue
		}
		res[i] = vowels[query[1]] - vowels[query[0]-1] // Check the previous index before query left range
	}
	return res
}

func isVowel(word string) bool {
	return isVowelChar(word[0]) && isVowelChar(word[len(word)-1])
}

func isVowelChar(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
