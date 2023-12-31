/*
@Time : 2021/7/8 15:29
@Author : sunxy
@File : solution
@description:
*/
package _30

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	cnt := map[int]int{}
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return
}
