/*
Package _2594_repairCars
@Time : 2023/9/7 10:04
@Author : sunxy
@File : repairCars
@description: https://leetcode.cn/problems/minimum-time-to-repair-cars/description/
*/
package _2594_repairCars

import "math"

func repairCars(ranks []int, cars int) int64 {
	// 二分法， 先取一个工人修理全部车的时间作为时间上限， 1为时间下限
	// 修理n 辆车， 时间是r*n^2   所以， t时间内可以修理 sqrt(t/r)
	// 计算所有人 t时间内修理的车辆总和，如果大于cars， 则上限改为t，否则下限为t+1
	var l, r int64
	l, r = 1, int64(ranks[0]*cars*cars)
	for l < r {
		t := (l + r) >> 1
		sum := 0
		for i := 0; i < len(ranks); i++ {
			sum += int(math.Floor(math.Sqrt(float64(t / int64(ranks[i])))))
		}
		if sum < cars {
			l = t + 1
		} else {
			r = t
		}
	}
	return r
}
