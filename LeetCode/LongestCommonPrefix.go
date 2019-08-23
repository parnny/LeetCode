package LeetCode

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	it := 0
	for {
		c := uint8(0)
		b := false
		for _,v:= range strs {
			if len(v) == it {
				b = false
				break
			}
			if c == 0 {
				c = v[it]
				b = true
			} else {
				b = b && ( v[it] == c )
			}
		}
		if !b {
			break
		}
		prefix += string(c)
		it++
	}
	return prefix
}