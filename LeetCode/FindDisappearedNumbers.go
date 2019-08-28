package LeetCode

import (
	"math"
)

func FindDisappearedNumbers(nums []int) []int {
	var result []int
	for _,v := range nums {
		if (v > 0) {
			nums[v-1] = -int(math.Abs(float64(nums[v-1])))
		} else {
			nums[-v-1] = -int(math.Abs(float64(nums[-v-1])))
		}
	}

	for i,v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

