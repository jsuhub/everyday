package dataStructure

// import (
// 	"fmt"
// )

// type node struct {
// 	row int
// 	col int
// 	val int
// }

// // func main() {
// // 	var cheaperMap [10][10]int
// // 	cheaperMap[1][2] = 1 //黑子为1
// // 	cheaperMap[2][3] = 2 //白字为2
// // 	for _, v := range cheaperMap {
// // 		for _, v2 := range v {
// // 			fmt.Print(v2, "\t")

// // 		}
// // 		fmt.Println()
// 	// }
// 	//转化稀疏数组
// 	//遍历数组 如果发现不为默认值 那么就创建node结构体放入切片中
// 	var parse []node
// 	var parseNode = node{
// 		row: 10,
// 		col: 10,
// 		val: 0,
// 	}
// 	parse = append(parse, parseNode)
// 	for i, v := range cheaperMap {
// 		for j, v2 := range v {
// 			if v2 != 0 {
// 				var parseNode = node{
// 					row: i,
// 					col: j,
// 					val: v2,
// 				}
// 				parse = append(parse, parseNode)

// 			}
// 		}
// 	}
// 	//输出稀疏数组
// 	for _, v := range parse {
// 		fmt.Print(v)
// 	}
// 	//存在一个文件中
// 	//恢复原始数组
// 	/*1.打开文件
// 	创建原始数组
// 	*/
// 	var chessMap [10][10]int

// 	for i, v := range parse {
// 		if i == 0 {
// 			continue
// 		}
// 		fmt.Print(v)
// 		chessMap[v.row][v.col] = v.val
// 		if i == len(parse) {
// 			break
// 		}
// 	}

// 	for _, v := range chessMap {
// 		for _, v2 := range v {
// 			fmt.Print(v2, "\t")

// 		}
// 		fmt.Println()
// 	}

// }

// //稀疏矩阵中的第一行表示矩阵的大小
// //且对数组定义时不能使用变量而是只能使用常量因此恢复从文件中捣毁只能运用切片去生成新的稀疏矩阵
