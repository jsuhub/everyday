package dataStructure //双向循环列表

import (
	"errors"
	"fmt"
)

/*
表示在头部/在尾部添加节点从而实现双向链表
*/
type Node struct {
	no    int
	name  string
	later *Node
	front *Node
}

func InsertNode(node *Node, nextnode *Node) error {
	if node != nil {
		nextnode.later = node.later
		nextnode.front = node
		node.later.front = nextnode
		node.later = nextnode
		return nil
	}
	return errors.New("数组为空不可添加")
}

func Listdouble(node *Node) error {
	if node == nil {
		return errors.New("数组为空")
	}
	head := node
	fmt.Println(node.name)
	node = node.later
	for node != head {
		fmt.Println(node.name)
		node = node.later
	}
	return nil
}

// n表示删除的是第几个
func Deldouble(node *Node, n int) (err error, del *Node) {
	if node == nil {
		return errors.New("数组为空"), nil
	}
	for i := 0; i < n; i++ {
		node = node.later
	}
	a := new(Node)
	a = node
	node.front.later = node.later
	node.later.front = node.front
	return nil, a
}
