package dataStructure

import "errors"

/*
规定从tail 和 head 都是从0开始 每次循环后
队列为满的时候是  (tail+1)%maxSize==head
*/
type CircleQueue struct {
	maxSize     int
	head        int
	tail        int
	circleArray [5]int
}

//判断是否为空的方法
func (this *CircleQueue) Empty() bool {
	if this.head%this.maxSize == this.tail%this.maxSize {
		return true
	}
	return false
}

//判断是否为满
func (this *CircleQueue) Full() bool {
	if (this.tail+1)%this.maxSize == this.head%this.maxSize {
		return true
	}
	return false
}

func (this *CircleQueue) Push(val int) error {
	if this.Full() {
		return errors.New("队列已满")
	}
	this.tail++
	this.circleArray[this.tail%this.maxSize] = val
	return nil
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.Empty() {
		return 0, errors.New("队列为空")
	}
	this.head++
	val = this.circleArray[this.head%this.maxSize]
	return val, nil

}

func (this *QueueArray) Size() int {
	size := (this.rear%this.maxSize - this.fornt%this.maxSize + this.maxSize) % this.maxSize
	return size
}
