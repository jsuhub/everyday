/*
1684. 统计一致字符串的数目
给你一个由不同字符组成的字符串 allowed 和一个字符串数组 words 。如果一个字符串的每一个字符都在 allowed 中，就称这个字符串是 一致字符串 。
请你返回 words 数组中 一致字符串 的数目。  ——————可以使用位运算解答

func main() {
    // 用位掩码表示字符集合
    allowed := "abc"
    mask := 0
    for _, char := range allowed {
        mask |= 1 << (char - 'a') // 将字符对应的位设置为 1
    }

    // 检查字符是否在集合中
    checkChar := func(char byte) bool {
        return mask&(1<<(char-'a')) != 0
    }

    fmt.Println(checkChar('a')) // true
    fmt.Println(checkChar('b')) // true
    fmt.Println(checkChar('c')) // true
    fmt.Println(checkChar('d')) // false
}
*/

package algorithm
import "fmt"
func countConsistentStrings(allowed string, words []string) int {
    number := 0
    allowedMap := make(map[rune]bool) //rune相当于char
    for _,v := range allowed {
      if !allowedMap[v]{
        allowedMap[v]=true 
      }  
    }
    fmt.Println(allowedMap)
    for _,word :=range words {
        sign := 0
        for _,v :=range word{
            if !allowedMap[v]{
                sign ++
                break 
            }
            }
        if(sign ==0){
                number ++
            }
    }
return number

}