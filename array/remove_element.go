package array

// RemoveElement https://leetcode.com/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return left
}
