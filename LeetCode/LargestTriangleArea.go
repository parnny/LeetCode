package LeetCode

import (
	"math"
)

func CalcTriangleAreaByEdge(a,b,c float64) float64 {
	p := (a+b+c)/2
	l := p*(p-a)*(p-b)*(p-c)
	s := math.Sqrt( l )
	return s
}

func CalcPointDistance(p1,p2 []int) float64 {
	x := float64(p1[0]-p2[0])
	y := float64(p1[1]-p2[1])
	return math.Sqrt(x*x + y*y)
}

func LargestTriangleArea(points [][]int) float64 {
	max := float64(-1)
	length := len(points)
	if  length < 3 {
		return 0
	}

	for x := 0; x < length-2; x++ {
		for y := x + 1; y < length-1; y++ {
			for z := y + 1; z < length; z++ {
				p1,p2,p3 := points[x],points[y],points[z]
				ds1 := CalcPointDistance(p1,p2)
				ds2 := CalcPointDistance(p2,p3)
				ds3 := CalcPointDistance(p3,p1)
				area := CalcTriangleAreaByEdge(ds1,ds2,ds3)
				if max < area {
					max = area
				}
			}
		}
	}

	return max
}