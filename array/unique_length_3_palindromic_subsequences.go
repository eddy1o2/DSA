package array

// CountPalindromicSubsequence https://leetcode.com/problems/unique-length-3-palindromic-subsequences/?envType=daily-question&envId=2025-01-04
func CountPalindromicSubsequence(s string) int {
	res := 0
	letters := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
	}
	for l, _ := range letters {
		j, i := 0, -1
		for k := 0; k < len(s); k++ {
			if s[k] == l {
				if i == -1 {
					i = k
				}
				j = k
			}
		}
		unique := make(map[byte]int)
		for k := i + 1; k < j; k++ {
			unique[s[k]]++
		}
		res += len(unique)
	}
	return res
}
