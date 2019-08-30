package LeetCode

import "strings"

func LengthOfLastWord(s string) int {
	result := strings.Split(s," ")

	for i := len(result) - 1; i >= 0; i-- {
		if len(result[i]) == 0 {
			continue
		}
		return len( result[i] )
	}

	return 0
}