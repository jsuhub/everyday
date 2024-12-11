package dataStructure

import (
	"errors"
	"fmt"
)

/*
队列是一个有序列表，可以用数组或链表实现
遵循先进先出的原则
数组实现队列就是定义一个头一个尾，但是当内部耗尽时就会实现没有利用到前面的————>这时就可以使用环形队列了
数组中得有maxSize    front rear 初始值都为-1
*/

/*
数组队列 添加元素的时候rear=rear+1
当front ==rear 时候表示队列为空
rear==maxSize 表示队列已经满了
*/
type QueueArray struct {
	maxSize    int
	fornt      int
	rear       int
	queuearray [3]int
}

func (this *QueueArray) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	}
	this.rear++
	this.queuearray[this.rear] = val
	return nil
}

func (this *QueueArray) ShowQueue() (err error) {
	if this.fornt > this.rear {
		return errors.New("队列存在错误")
	} else if this.fornt == this.rear {
		return errors.New("队列为空")
	}
	for {
		if this.fornt > this.rear {
			break
		}
		this.fornt++
		fmt.Println(this.queuearray[this.fornt])
	}
	return nil

}

func (this *QueueArray) GetQUeue() (val int, err error) {
	if this.fornt >= this.rear {
		return 0, errors.New("空间不足")
	}
	this.fornt++
	return this.queuearray[this.fornt], nil
}
