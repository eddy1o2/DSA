package array

import "strings"

// ShiftingLetters https://leetcode.com/problems/shifting-letters-ii/?envType=daily-question&envId=2025-01-05
func ShiftingLetters(s string, shifts [][]int) string {
	tmp := make([]int, len(s))
	for _, shift := range shifts {
		if shift[2] == 1 { //Forward
			tmp[shift[0]]++
			if shift[1] != len(s)-1 {
				tmp[shift[1]+1]--
			}
		} else {
			tmp[shift[0]]--
			if shift[1] != len(s)-1 {
				tmp[shift[1]+1]++
			}
		}
	}
	prefixSum := make([]int, len(tmp))
	total := 0
	for i := range tmp {
		total += tmp[i]
		prefixSum[i] = total
	}
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		changed := prefixSum[i]
		if prefixSum[i] >= 26 || prefixSum[i] <= -26 {
			changed = changed % 26
		}
		b := s[i] + byte(changed)
		if b > 122 {
			b = 96 + (b - 122) // backward wrapping
		} else if b < 97 {
			b = 123 - (97 - b) // forward wrapping
		}
		builder.WriteByte(b)
	}
	return builder.String()
}
