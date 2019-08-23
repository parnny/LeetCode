package LeetCode


func RomanToInt(s string) int {
	var mapping = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	this, next := "", ""
	end := len(s)-1
	for i,c := range s {
		this = string(c)
		res += mapping[this]
		if i!=end {
			next = string(s[i+1])
			if this == "I" && (next == "V" || next == "X") {
				res -= 2
			} else if this == "X" && (next == "L" || next == "C") {
				res -= 20
			} else if this == "C" && (next == "D" || next == "M") {
				res -= 200
			}
		}
	}

	return res
}