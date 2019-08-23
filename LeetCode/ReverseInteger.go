package LeetCode

import (
	"math"
)

func Reverse(x int) int {
	f := false
	r := 0
	if x < 0 {
		f = true
		x *= -1
	}

	fx := float64(x)
	for {
		r = r *10 + int(math.Mod(fx,10))
		fx /= 10
		if fx < 1 {
			break
		}
	}

	if r > math.MaxInt32 {
		return 0
	}
	if f {
		r *=-1
	}
	return r
}