package LeetCode

import (
	"math"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	fx := float64(x)
	r := 0
	for {
		r = r *10 + int(math.Mod(fx,10))
		fx /= 10
		if fx < 1 {
			break
		}
	}

	return (r == x)
}