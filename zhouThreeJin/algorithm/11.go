/*远征队即将开启未知的冒险之旅，不过在此之前，将对补给车队进行最后的检查。supplies[i] 表示编号为 i 的补给马车装载的物资数量。 考虑到车队过长容易被野兽偷袭，他们决定将车队的长度变为原来的一半（向下取整），计划为：
找出车队中 物资之和最小 两辆 相邻 马车，将它们车辆的物资整合为一辆。若存在多组物资之和相同的马车，则取编号最小的两辆马车进行整合；
重复上述操作直到车队长度符合要求。
请返回车队长度符合要求后，物资的分布情况。
*/
//如果按照重合一次再重新排练时间复杂度为O(n^2)
//思路：可以把其他的和保存住，然后每次修改就可以不需要再加加的动作了。可以想象一下最小堆的运算思路。把新的数字放到 做不出来无法判断更新是否正确
package algorithm

import (
	"container/heap"
	"fmt"
)

type SupplyNode struct {
	sum   int
	index int
}

type MinHeap []SupplyNode

// 对堆的重写确定是一个最小堆的要求
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum <= h[j].sum || h[i].index < h[j].index }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(SupplyNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func supplyWagon(supplies []int) []int {

	h := &MinHeap{}
	heap.Init(h) //初始化堆
	length := len(supplies)
	for i := 0; i < length-1; i++ {
		heap.Push(h, SupplyNode{supplies[i] + supplies[i+1], i}) //放入堆中并且自动生成最小堆
	}
	for len(supplies) > length/2 {
		fmt.Println(h)
		fmt.Println(supplies)

		Hnode := heap.Pop(h).(SupplyNode)                               //从最小堆里面提出最小值并且将自动再次排序。
		if supplies[Hnode.index]+supplies[Hnode.index+1] == Hnode.sum { //判断是否是无效堆结点————因为没有从堆里面删除结点所以存在无效结点(就是哪些没有实时更新的结点)。
			supplies[Hnode.index] = Hnode.sum
			supplies = append(supplies[:Hnode.index+1], supplies[Hnode.index+2:]...) //将最初的切片合并，并且从里面取出Hnode.index 这个值

			if Hnode.index > 0 {
				heap.Push(h, SupplyNode{supplies[Hnode.index] + supplies[Hnode.index-1], Hnode.index - 1})
			} //把更新后的堆结点插入最小堆中

			if Hnode.index < len(supplies)-1 {
				heap.Push(h, SupplyNode{supplies[Hnode.index] + supplies[Hnode.index+1], Hnode.index})

			}
		}
	}

	return supplies
}
