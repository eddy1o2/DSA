package array

import "strings"

// CountPrefixSuffixPairs https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/?envType=daily-question&envId=2025-01-08
func CountPrefixSuffixPairs(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				res++
			}
		}
	}
	return res
}

func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}
