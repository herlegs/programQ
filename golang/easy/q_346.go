package easy

import "container/list"

type MovingAverage struct {
	size int
	list *list.List
	sum int
	curSize int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		list: list.New(),
	}
}


func (this *MovingAverage) Next(val int) float64 {
	this.list.PushBack(val)
	this.sum += val
	this.curSize++
	if this.list.Len() > this.size {
		front := this.list.Front()
		this.list.Remove(front)
		this.sum -= front.Value.(int)
		this.curSize--
	}
	return float64(this.sum) / float64(this.curSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
