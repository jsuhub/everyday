package algorithm

// import (
// 	"sort"
// )

// type colArray struct {
// 	Col   int
// 	Value int64
// 	Row   int
// }
// type ByValue []colArray

// func (v ByValue) Len() int           { return len(v) }
// func (v ByValue) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
// func (v ByValue) Less(i, j int) bool { return v[i].Value < v[j].Value }
// func maxSpending(values [][]int) int64 {
// 	row := len(values)
// 	col := len(values[0])
// 	rowArray := make([]colArray, row)
// 	for i := 0; i < row; i++ {
// 		rowArray[i].Value = int64(values[i][col-1])
// 		rowArray[i].Col = col - 1
// 		rowArray[i].Row = i
// 	}
// 	sort.Sort(ByValue(rowArray))
// 	var sum, day int64
// 	sum = 0
// 	day = 1
// 	for day <= int64(col*row) {
// 		sum = sum + rowArray[0].Value*day
// 		rowArray[0].Col--
// 		if rowArray[0].Col == -1 {
// 			rowArray = rowArray[1:]
// 			if len(rowArray) == 1 {
// 				for {
// 					day++
// 					sum = sum + day*rowArray[0].Value
// 					rowArray[0].Col--
// 					if rowArray[0].Col == -1 {
// 						break
// 					}
// 					rowArray[0].Value = int64(values[rowArray[0].Row][rowArray[0].Col])
// 				}
// 				break
// 			}
// 			day++
// 			continue
// 		}
// 		rowArray[0].Value = int64(values[rowArray[0].Row][rowArray[0].Col])
// 		indexValue := rowArray[0].Value
// 		for i := 1; i < len(rowArray) && indexValue > rowArray[i].Value; i++ {
// 			rowArray[i-1], rowArray[i] = rowArray[i], rowArray[i-1]
// 		}
// 		day++
// 	}
// 	return int64(sum)
// }
