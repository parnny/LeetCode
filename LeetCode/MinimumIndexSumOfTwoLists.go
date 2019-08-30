package LeetCode

var FindRestaurantInput1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
var FindRestaurantInput2 = []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

func FindRestaurant(list1 []string, list2 []string) []string {
	record := make(map[string]int)

	for i,v := range list1 {
		record[v] = i
	}

	min := len(list1) + len(list2)
	flag := 0
	var result []string
	for i,v := range list2 {
		if j, ok := record[v]; ok {
			flag = i + j
			if flag < min {
				result = []string{v,}
				min = flag
			} else if flag == min {
				result = append(result, v)
			}
		}
	}
	return result
}