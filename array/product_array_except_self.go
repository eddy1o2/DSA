package array

// ProductExceptSelf https://leetcode.com/problems/product-of-array-except-self/
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	totalPrefix := 1
	totalSuffix := 1
	res[0] = totalPrefix
	for i := 1; i < len(nums); i++ {
		totalPrefix *= nums[i-1]
		res[i] = totalPrefix
	}
	res[len(nums)-1] = totalSuffix * res[len(nums)-1]
	for j := len(nums) - 2; j >= 0; j-- {
		totalSuffix *= nums[j+1]
		res[j] = totalSuffix * res[j]
	}
	return res
}
