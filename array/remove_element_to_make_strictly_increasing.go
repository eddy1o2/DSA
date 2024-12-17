package array

// CanBeIncreasing https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
func CanBeIncreasing(nums []int) bool {
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] >= nums[i+1] {
			previous := i
			if i > 0 && nums[i+1] <= nums[i-1] {
				previous = i + 1
			}
			nums = remove(nums, previous)
			break
		}
	}
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
