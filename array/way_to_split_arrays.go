package array

// WaysToSplitArray https://leetcode.com/problems/number-of-ways-to-split-array/?envType=daily-question&envId=2025-01-03
func WaysToSplitArray(nums []int) int {
	prefixSum := make([]int, len(nums))
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSum[i] = sum
	}
	for i := 0; i < len(nums)-1; i++ {
		if prefixSum[i] >= sum-prefixSum[i] {
			res++
		}
	}
	return res
}
