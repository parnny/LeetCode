package LeetCode

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsPrime(n int32) bool {
	if n <= 3 {
		return n > 1
	}

	t := int32(math.Mod(float64(n),6))
	if t != 1 && t != 5	 {
		return false
	}

	q := int32(math.Sqrt(float64(n)))
	for i := int32(5); i <= q; i += 6 {
		if math.Mod(float64(n), float64(i)) == 0 || math.Mod(float64(n), float64(i+2)) == 0 {
			return false
		}
	}
	return true
}