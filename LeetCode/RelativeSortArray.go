package LeetCode

import "sort"

type list []int

func (a list) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var record map[int]int

func (a list)Len() int {
	return len(a)
}

func (a list)Less(i,j int) bool {
	v1,v2 := a[i],a[j]
	w1,ok1 := record[ v1 ]
	w2,ok2 := record[ v2 ]
	if !ok1 && !ok2 {
		return v1 < v2
	}
	return w1 > w2
}

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	maxW := len(arr2)
	record = make(map[int]int)
	for i,v := range arr2 {
		record[v] = maxW - i
	}
	var a list
	a = arr1
	sort.Sort(a)
	return a
}