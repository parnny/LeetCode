package LeetCode

import "math"

func isSelfDividingNumbers(num float64) bool {
	x := num
	for {
		y := math.Mod(x,10)
		if math.Mod(num, y) != 0 {
			return false
		}
		x = (x - y)/10
		if x < 1 {
			break
		}
	}
	return true
}

func SelfDividingNumbers(left int, right int) []int {
		var result []int
		for i := left; i <= right; i++ {
			if isSelfDividingNumbers(float64(i)) {
				result = append(result, i)
			}
		}
		return result
}