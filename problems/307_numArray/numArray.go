/*
Package _07_numArray
@Time : 2023/11/13 11:44
@Author : sunxy
@File : numArray
@description:
*/
package _07_numArray

import "fmt"

type NumArray struct {
	nums     []int
	retCache map[string]int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums:     nums,
		retCache: make(map[string]int),
	}
}

func (this *NumArray) Update(index int, val int) {
	this.nums[index] = val

}

func (this *NumArray) SumRange(left int, right int) int {
	key := fmt.Sprintf("%d:%d", left, right)
	if v, ok := this.retCache[key]; ok {
		return v
	}
	ret := 0
	for i := left; i <= right; i++ {
		ret += this.nums[i]
	}
	this.retCache[key] = ret
	return ret
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
