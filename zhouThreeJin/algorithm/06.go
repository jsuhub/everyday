package algorithm

import "fmt"

func minValidStrings(words []string, target string) int {
	db := make([]int, len(target)+1)
	db[0] = 1
	per := make(map[string]bool)
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			per[word[:i]] = true
		}
		fmt.Println(per)
		j := 0
		for i := 1; i <= len(target); i++ {
			k := 0
			for per[target[j:i]] != true {
				db[i] = db[i-1] + 1
				j = i
				i = i + 1
				k++
				if k == 2 {
					return -1
				}
			}
			db[i] = db[i-1]
		}
	}
	return db[len(target)+1]
}
