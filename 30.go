package main

import (
	"container/list"
)

type MinStack struct {
	A *list.List
	B *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		A: list.New(),
		B: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.A.PushBack(x)

	b := this.B.Back()
	if b == nil || b.Value.(int) >= x {
		this.B.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	a := this.A.Back()
	b := this.B.Back()

	if a == nil {
		return
	}

	this.A.Remove(a)

	if b == nil {
		return
	}
	if a.Value.(int) == b.Value.(int) {
		this.B.Remove(b)
	}

}

func (this *MinStack) Top() int {
	return this.A.Back().Value.(int)
}

func (this *MinStack) Min() int {
	b := this.B.Back()
	if b != nil {
		return b.Value.(int)
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
