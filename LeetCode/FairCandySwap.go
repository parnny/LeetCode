package LeetCode

func FairCandySwap(A []int, B []int) []int {
	SumA, SumB := 0,0
	MapA := map[int]bool{}
	MapB := map[int]bool{}

	res := []int{}

	for _,v := range A {
		SumA += v
		MapA[v] = true
	}
	for _,v := range B {
		SumB += v
		MapB[v] = true
	}

	diff := (SumB - SumA)/2
	for _,v := range A {
		if _,ok := MapB[v+diff]; ok {
			res = append(res, v, v+diff)
			break
		}
	}

	return res
}