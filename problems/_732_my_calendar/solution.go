/*
Package _732_my_calendar
@Time : 2022/6/6 15:39
@Author : sunxy
@File : solution
@description:
*/
package _732_my_calendar

type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (t MyCalendarThree) add(x, delta int) {
	if val, ok := t.Get(x); ok {
		delta += val.(int)
	}
	t.Put(x, delta)
}

func (t MyCalendarThree) Book(start, end int) (ans int) {
	t.add(start, 1)
	t.add(end, -1)

	maxBook := 0
	for it := t.Iterator(); it.Next(); {
		maxBook += it.Value().(int)
		if maxBook > ans {
			ans = maxBook
		}
	}
	return
}
