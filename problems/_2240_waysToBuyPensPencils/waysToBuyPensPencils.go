/*
Package _2240_waysToBuyPensPencils
@Time : 2023/9/1 17:40
@Author : sunxy
@File : waysToBuyPensPencils
@description: https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/
*/
package _2240_waysToBuyPensPencils

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var res int64 = 0
	g := total / cost1
	for i := 0; i <= g; i++ {
		res += 1
		q := (total - cost1*i) / cost2
		if q > 0 {
			res += int64(q)
		} else {
			continue
		}
	}
	return res
}
