package LeetCode

func CountPrimeSetBits(L int, R int) int {
	result := 0
	for v := L; v <= R; v++ {
		c := HammingWeight(uint32(v))

		if IsPrime(int32(c)) {
			result++
		}
	}
	return result
}
