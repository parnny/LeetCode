package LeetCode

import (
	"strconv"
)

var BaseballGameCalPointsInput1 = []string{"5","2","C","D","+"}
var BaseballGameCalPointsInput2 = []string{"5","-2","4","C","D","9","+","+"}

func BaseballGameCalPoints(ops []string) int {
	var valid []int
	sum := 0
	v := 0
	var err error
	for _,op := range ops {
		switch op {
		case "C":
			sum-=valid[len(valid)-1]
			valid = valid[:len(valid)-1]
			break
		case "D":
			v = valid[len(valid)-1]*2
			sum+=v
			valid = append(valid, v)
			break
		case "+":
			if len(valid) > 1 {
				v = valid[len(valid)-1] + valid[len(valid)-2]
				sum+=v
				valid = append(valid, v)
			}
		default:
			v,err =strconv.Atoi(op)
			if err == nil {
				sum+=v
			}
			valid = append(valid, v)
			break
		}
	}
	return sum
}