// import "fmt"
// func longestCommonPrefix(strs []string) string {
//  if strs[0]== "" {
//  return ""
//  }
//  words := make([]byte,0) 
//  j:=0
//  word :=strs[0][0]
// for {
//     fmt.Println(string(word))
//     for i:=0;i<len(strs);i++{
//         if( j==len(strs[i]) || word != strs[i][j]) {
//             return string(words)
//         }
//     }
//     words = append(words,word)
//     j++
//     if j==len(strs[0]){
//         return string(words)
//     }
//     word=strs[0][j]
// }
// }