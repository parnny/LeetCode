package LeetCode

import (
	"sort"
)

func DeckRevealedIncreasing(deck []int) []int {
	l := len(deck)
	var src []int
	var sim []int
	var res []int
	for i:=0; i<l; i++ {
		src = append(src, i)
		res = append(res, i)
	}
	for len(src) != 0 {
		first := src[0]
		src = src[1:]
		sim = append(sim, first)
		if len(src) != 0 {
			second := src[0]
			src = src[1:]
			src = append(src, second)
		}
	}
	sort.Ints(deck)
	for i,v := range sim {
		res[v] = deck[i]
	}
	return res
}