package algorithm

// import "sort"

// func getNext(str string) (next []int) {
// 	j := 0
// 	next = make([]int, len(str))
// 	for i := 1; i < len(str); i++ {
// 		for j > 0 && str[i] != str[j] {
// 			j = next[j-1]
// 		}
// 		if str[i] == str[j] {
// 			j++
// 		}
// 		next[i] = j
// 	}
// 	return next
// }

// func minValidStrings(words []string, target string) int {
// 	wordss := words
// 	sort.Slice( //先按照字符串长度排序
// 		wordss, func(i, j int) bool {
// 			return len(wordss[i]) < len(wordss[j])
// 		})
// 	for _, v := range wordss {
// 		j := 0
// 		next := getNext(v)
// 		for i := 0; i < len(target); i++ { //每个使用kmp算法作比较
// 			for j > 0 && v[j] != target[i] {
// 				j = next[j-1]
// 			}
// 			if v[j] == target[i] {
// 				j++
// 			}
// 			if j == len(v) {
// 				return len(v)
// 			}
// 		}
// 	}
// 	return -1
// }
