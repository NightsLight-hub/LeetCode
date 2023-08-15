/*
Package _2208_halveArray
@Time : 2023/7/25 9:44
@Author : sunxy
@File : halveArray
@description:
*/
package _2208_halveArray

import "github.com/emirpasic/gods/queues/priorityqueue"

func byPriority(a, b interface{}) int {
	if a.(float64) > b.(float64) {
		return -1
	}
	return 1
}

func halveArray(nums []int) int {
	var sum float64 = 0
	queue := priorityqueue.NewWith(byPriority)
	for _, v := range nums {
		queue.Enqueue(float64(v))
		sum += float64(v)
	}
	sum /= 2
	res := 0
	for {
		e, ok := queue.Dequeue()
		if !ok {
			return -1
		}
		t := e.(float64) / 2
		sum -= t
		res++
		if sum <= 0 {
			return res
		}
		queue.Enqueue(t)
	}
}
