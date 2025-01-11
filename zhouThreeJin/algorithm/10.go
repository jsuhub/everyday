/*题目：
给定两个数组 nums1 和 nums2 ，返回它们的交集
输出结果中的每个元素一定是 唯一 的。我们可以不考虑输出结果的顺序。
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
//思路： 1.从两个数组里面选择一个去掉其中的重复的值，再用另一个的每个都去搜索一篇  2.使用Map类型去存储数据就可以直接调用了
package algorithm

func intersection(nums1 []int, nums2 []int) []int {
	var numMap = make(map[int]int)
	for _, v := range nums1 {
		numMap[v] = 1
	}
	var interSection []int = nil
	for _, v := range nums2 {
		if numMap[v] == 1 {
			numMap[v]++
			interSection = append(interSection, v)

		}
	}
	return interSection
}
