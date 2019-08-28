package LeetCode

import "math"

func MajorityElement(nums []int) int {
	limit := int(math.Ceil( float64(len(nums))/2 ))
	record := make(map[int]int)	// <value, counting>
	for _,v := range nums {
		if _, ok := record[v]; !ok {
			record[v] = 1
		} else {
			record[v]++
		}
		if record[v] >= limit {
			return v
		}
	}
	return 0
}