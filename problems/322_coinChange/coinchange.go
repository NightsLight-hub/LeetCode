package _22_coinChange

import "math"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/coin-change
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// amount < min(coins)   dp(coins, amount ) = -1
// amount = 0            dp(coins, amount ) = 0
// min(dp(coins, amount - coin)) + 1, coin 属于 coins, if all dp(coins, amount - coin) <0   return -1

var m map[int]int
var bigValue int = math.MaxInt32 - 1

func coinChange(coins []int, amount int) int {
	m = make(map[int]int)
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	res := bigValue
	for _, coin := range coins {
		subDb := 0
		if value, ok := m[amount-coin]; ok {
			subDb = value
		} else {
			subDb = dp(coins, amount-coin)
		}
		if subDb < 0 {
			continue
		}
		res = min(res, subDb+1)
	}
	if res == bigValue {
		res = -1
	}
	m[amount] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
