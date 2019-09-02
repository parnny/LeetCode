package LeetCode

import "sort"

func LargestPerimeterTriangle(A []int) int {
	sort.Ints(A)
	max := 0
	for i := len(A)-1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i]+A[i-1]+A[i-2]
		}
	}
	return max
}