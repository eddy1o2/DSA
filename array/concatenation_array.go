package array

// GetConcatenation https://leetcode.com/problems/concatenation-of-array/
func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i, _ := range ans {
		if i >= len(nums) {
			ans[i] = ans[i-len(nums)]
			continue
		}
		ans[i] = nums[i]
	}
	return ans
}
