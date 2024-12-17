package algorithm

// /*KMP算法中主串指针没有必要回退
// 因为我们已经匹配了前面的相同的部分了所以不需要再返回主串的只需要确认此时的next数组里面对应的值i然后模式串里面的第i个字符再和主串不匹配的比较重复这个步骤直到模式串里面没有可以匹配的就下一个，重新匹配 ————这样主串就不需要返回了
// 难点在于求next数组
// 这个时间复杂度为O(n+m)
// 首先前缀数组里面的下标和字串的下标是一一对应的
// goole这个字串中next[0]不占且总共到next[6]
// 当第一个不匹配的时候next[1]=0 表示重新主串中的第i+1个和子串的s[1]相匹配
// next[2]=1表示当主串和子串第二个不匹配的时候重新跳转到主串的第一个再继续比较
// 从第三个开始时，当不匹配的时候就可以划一条线
// go|cdsa		发现左边都一样的可以滑动这个下面的部分     go|cdsa
// go|ole											 	  |goole	此时无法匹配所以直接写1 next[3]=1	 并且全部移动到模式串的右边

// 假如匹配的上
// i
// goog|aaa   发现有匹配不上的于是右移    goog|aa		 所以此时只需要next[5]=2表示下次只需要比较第二位是否匹配的上
// j
// goog|le									g|oogle						//快速手算的方法 假设前面相等然后在不相等前面画一条竖线 然后移动子串进行比较有几个最大相同的就是写几
// */
// func buildPrefixTable(needle string) []int {
// 	m := len(needle)
// 	prefixTable := make([]int, m)
// 	j := 0

// 	for i := 1; i < m; i++ {
// 		for j > 0 && needle[i] != needle[j] {
// 			j = prefixTable[j-1]
// 		}
// 		if needle[i] == needle[j] {
// 			j++
// 		}
// 		prefixTable[i] = j
// 	}

// 	return prefixTable
// }

// // KMP 算法实现查找第一个匹配项的下标
// func strStr(haystack string, needle string) int {
// 	if len(needle) == 0 {
// 		return 0 // 如果 needle 为空字符串，返回 0
// 	}
// 	m, n := len(haystack), len(needle)
// 	prefixTable := buildPrefixTable(needle)
// 	j := 0
// 	for i := 0; i < m; i++ {
// 		// 匹配失败时，根据部分匹配表跳过不必要的字符
// 		for j > 0 && haystack[i] != needle[j] {
// 			j = prefixTable[j-1]
// 		}
// 		if haystack[i] == needle[j] {
// 			j++
// 		}
// 		// 如果整个 needle 匹配完成，返回匹配的起始位置
// 		if j == n {
// 			return i - n + 1
// 		}
// 	}
// 	return -1 // 如果没有找到匹配项，返回 -1
// }

// //前缀表是从0开始的，
