/*
Package _2651_findDelayedArrivalTime
@Time : 2023/9/8 16:01
@Author : sunxy
@File : findDelayedArrivalTime
@description: https://leetcode.cn/problems/calculate-delayed-arrival-time/
*/
package _2651_findDelayedArrivalTime

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
