package LeetCode

func TwoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i := 0; i< len(nums); i++ {
		v,ok := mapping[ target-nums[i] ]
		if ok {
			return []int{v,i}
		}
		mapping[ nums[i] ] = i
	}
	return []int{0,0}
}