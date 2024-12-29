package algorithm

func minValidStrings(words []string, target string) int {
	db := make([]int, len(target)+1)
	db[0] = 0
	var per map[string]bool
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			per[word[:i]] = true
		}
		j := 0
		k := 0
		for i := 1; i < len(target); i++ {

			if per[target[j:i]] != true {
				db[i] = db[i-1] + 1
				j = i - 1
				k++
			}

			db[i] = db[i-1]
		}
	}
	return db[len(target)]
}
