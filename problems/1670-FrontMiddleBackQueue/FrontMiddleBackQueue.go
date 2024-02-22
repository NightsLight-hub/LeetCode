/*
Package _670_FrontMiddleBackQueue
@Time : 2023/11/28 14:29
@Author : sunxy
@File : FrontMiddleBackQueue
@description: https://leetcode.cn/problems/design-front-middle-back-queue/?envType=daily-question&envId=Invalid%20Date
*/
package _670_FrontMiddleBackQueue

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

type FrontMiddleBackQueue struct {
	head   *ListNode
	middle *ListNode
	tail   *ListNode
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		head:   nil,
		middle: nil,
		tail:   nil,
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	if this.head == nil {
		this.head = &ListNode{
			val:  val,
			prev: nil,
			next: nil,
		}
		this.middle = this.head
		this.tail = this.head
	} else {
		node := &ListNode{
			val:  val,
			prev: nil,
			next: this.head,
		}
		this.head.prev = node
		this.head = node
		if this.head.next == this.middle {
			this.middle = this.head
		}
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {

}

func (this *FrontMiddleBackQueue) PushBack(val int) {

}

func (this *FrontMiddleBackQueue) PopFront() int {

}

func (this *FrontMiddleBackQueue) PopMiddle() int {

}

func (this *FrontMiddleBackQueue) PopBack() int {

}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
