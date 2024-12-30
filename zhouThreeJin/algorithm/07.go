// 提供一种思路可以直接在再生成一个一摸一样的🌲不过该树的结点表示的是子节点之和即左结点加右结点再加上本身就可以让下面代码执行左右结点差的时候不需要每次重复遍历大大减少了时间且只增加了一个树的内存大小。
package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := getAll(root.Right)
	left := getAll(root.Left)
	var sum int
	if right >= left {
		sum = right - left
	} else {
		sum = left - right
	}
	sum = sum + findTilt(root.Left) + findTilt(root.Right)
	return sum
}

func getAll(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getAll(root.Left) + getAll(root.Right) + root.Val

}
