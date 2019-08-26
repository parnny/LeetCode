package LeetCode

import (
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	counting := make(map[string]int)

	for _,domain := range cpdomains {
		strPair := strings.Split(domain," ")
		count, _ := strconv.Atoi(strPair[0])
		splitDomain := strings.Split(strPair[1],".")
		for i := 0; i < len(splitDomain); i++ {
			k := strings.Join(splitDomain[i:],".")
			if _, ok := counting[k]; !ok {
				counting[k] = count
			} else {
				counting[k]+=count
			}

		}
	}

	var result []string
	for s,n := range counting {
		result = append(result, strings.Join([]string{strconv.Itoa(n), string(s)}," ") )
	}
	return result
}