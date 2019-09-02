package LeetCode

func FindComplement(num int) int {
	v := 1
	for {
		if v > num {
			break
		}
		v = v << 1
	}
	return num ^ (v-1)
}

