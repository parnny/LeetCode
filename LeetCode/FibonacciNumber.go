package LeetCode

func FibonacciNumber(N int) int {
	if N <= 1 {
		return N
	}
	return FibonacciNumber(N-1) + FibonacciNumber(N-2)
}