// import "math"
package algorithm

// func thirdMax(nums []int) int {
// 	var max int
// 	key := 1
// 	max = nums[0]
// 	max2 := math.MinInt
// 	max3 := math.MinInt
// 	for _, v := range nums {
// 		if v >= max {
// 			if v == max {
// 				continue
// 			}
// 			max3 = max2
// 			max2 = max
// 			max = v
// 			key++
// 			continue
// 		} else if v >= max2 {
// 			if v == max2 {
// 				continue
// 			}
// 			max3 = max2
// 			max2 = v
// 			key++
// 			continue
// 		} else if v >= max3 {
// 			if v == max3 {
// 				continue
// 			}
// 			max3 = v
// 			key++
// 		}
// 	}
// 	if key >= 3 {
// 		return max3
// 	}
// 	return max
// }
