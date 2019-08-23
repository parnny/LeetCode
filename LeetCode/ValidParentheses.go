package LeetCode


func IsValidParentheses(s string) bool {
	var stack []int32

	right := map[int32]int32 {
		')': '(',
		']': '[',
		'}': '{',
	}

	for _,c := range s {
		if l, ok := right[c]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == l {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}