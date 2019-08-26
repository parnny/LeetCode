package LeetCode

func TotalHammingDistance(nums []int) int {
	out := 0
	lengh := len(nums)
	max := 0
	for i := uint(0); i < 32; i++ {
		c, b := 0, 1<<i
		if i != 0 && b > max {
			break
		}
		for _, num := range nums {
			if i == 0 {
				if max < num {
					max = num
				}
			}
			if num&b != 0 {
				c++
			}
		}
		out += c * (lengh - c)
	}
	return out
}