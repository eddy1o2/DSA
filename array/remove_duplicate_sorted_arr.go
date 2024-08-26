package array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array

func RemoveDuplicates(nums []int) int {
	left, right := 1, 1
	for right < len(nums) {
		if nums[right] == nums[right-1] {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}
	return left
}
