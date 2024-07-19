/*
Package quicksort
@Time : 2023/7/28 15:00
@Author : sunxy
@File : quickSort2
@description:
*/
package quicksort

func quickSort2[T any](nums []T, f func(a, b T) bool) {
	if len(nums) <= 1 {
		return
	}
	// mid 是这一轮选的基准值，本轮的目的就是把基准值放到最终有序的位置上
	mid := nums[0]
	midIndex, tailIndex := 0, len(nums)-1
	// 将小的都放在基准值左边，大的放在基准值右边，右边从尾巴开始放，左边从头开始放
	// 因为基准值在头部本来就占有一个位置，可以认为是一个活动的坑，因此给了利用原数组就可以排序的空间
	// 其实如果引入另外一个数组，排序复杂度能降低很多
	for i := 1; i <= tailIndex; {
		if f(nums[i], mid) {
			// 如果小于基准值，就把这个位置的数值与基准值调换
			nums[i], nums[midIndex] = nums[midIndex], nums[i]
			// 基准值的索引更新
			midIndex = i
			i++
		} else {
			// 大于基准值，就把这个位置的数值与tail调换
			nums[i], nums[tailIndex] = nums[tailIndex], nums[i]
			// tail 往前一位
			tailIndex--
		}
	}
	quickSort2(nums[:midIndex], f)
	quickSort2(nums[midIndex+1:], f)
}
