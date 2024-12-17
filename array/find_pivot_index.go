package array

// PivotIndex https://leetcode.com/problems/find-pivot-index/
func PivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left := 0
		right := 0
		left = sumArr(nums[:i])
		// if i==1 {
		// 	left = sumArr(nums[:i])
		// } else {
		// 	left = sumArr(nums[:i-1])
		// }
		right = sumArr(nums[i+1:])
		if left == right {
			return i
		}
	}
	return -1
}

func sumArr(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}
