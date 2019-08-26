package LeetCode

func IsPowerOfTwo(n int) bool {
	return n & (n-1) == 0
}