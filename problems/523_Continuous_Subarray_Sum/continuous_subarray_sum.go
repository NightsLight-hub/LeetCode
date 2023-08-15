/*
@Time : 2021/6/2 15:35
@Author : sunxy
@File : continuous_subarray_sum
@description:
*/
package _523_Continuous_Subarray_Sum

/**
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。

如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/continuous-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkSubarraySum(nums []int, k int) bool {
	m := len(nums)
	if m < 2 {
		return false
	}
	mp := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if prevIndex, has := mp[remainder]; has {
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}
