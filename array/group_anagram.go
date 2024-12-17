package array

// GroupAnagrams https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	tmp := make(map[[26]int][]string)
	for _, str := range strs {
		anagramHash := [26]int{}
		for i := 0; i < len(str); i++ {
			anagramHash[str[i]-'a']++
		}
		tmp[anagramHash] = append(tmp[anagramHash], str)
	}
	result := make([][]string, 0, len(tmp))
	for _, group := range tmp {
		result = append(result, group)
	}
	return result
}
