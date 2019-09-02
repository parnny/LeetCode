package LeetCode

func PeakIndexInMountainArray(A []int) int {
	for i,v := range A {
		if i > 0 && v < A[i-1]  {
			return i-1
		}
	}
	return 0
}