package LeetCode

import "math"

func SortArrayByParityII(A []int) []int {
	result := make([]int,len(A))
	p1,p2 := 0, 1
	for _,v := range A {
		if int(math.Mod( float64(v), float64(2) )) == 0 {
			result[p1] = v
			p1+=2
		} else {
			result[p2] = v
			p2+=2
		}
	}
	return result
}