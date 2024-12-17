package array

// ContainsDuplicate https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	checker := make(map[int]bool)
	for _, num := range nums {
		if _, ok := checker[num]; ok {
			return true
		}
		checker[num] = true
	}
	return false
}
