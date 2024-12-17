package array

// TwoSum https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, num := range nums {
		if v, ok := tmp[num]; ok {
			return []int{v, i}
		}
		tmp[target-num] = i
	}
	return []int{}
}
