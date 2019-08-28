package LeetCode

func HammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		if num & 1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}