package array

// TopKFrequent https://leetcode.com/problems/top-k-frequent-elements/
func TopKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	buckets := make([][]int, len(nums)+1)
	for _, num := range nums {
		frequency[num]++
	}
	for num, f := range frequency {
		buckets[f] = append(buckets[f], num)
	}
	var result []int
	for i := len(buckets) - 1; i > 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
