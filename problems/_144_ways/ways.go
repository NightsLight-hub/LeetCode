/*
Package _144_ways
@Time : 2023/8/17 9:18
@Author : sunxy
@File : ways
@description:
*/
package _144_ways

import "math/big"

/*
*
问题可以简化为寻找苹果之间的连接点的位置。其实就是用刀在连接点出分割pizza。
由于刀只能垂直或者水平，所以寻找的方案如下：
1. 遍历所有列，当前列上的苹果，下表为 ai， 则连接点共有  (A1 - A0) + (A2-A1)+ .. + (Ai - A(i-1))
2. 遍历行， 连接点 计算法同上
问题退化为 从 n个连接点，寻找k-1个的组合问题
n!/((k-1)! * (n-k+1)!)
*/
func ways(pizza []string, k int) int {
	yl := len(pizza)
	xl := len(pizza[0])
	n := 0
	for y := 0; y < yl; y++ {
		var tmp int = -1
		for x := 0; x < xl; x++ {
			if pizza[y][x] == 'A' {
				if tmp != -1 {
					n += x - tmp
				}
				tmp = x
			}
		}
	}
	for x := 0; x < xl; x++ {
		var tmp int = -1
		for y := 0; y < yl; y++ {
			if pizza[y][x] == 'A' {
				if tmp != -1 {
					n += y - tmp
				}
				tmp = y
			}
		}
	}
	f := func(num int) *big.Int {
		var res = big.NewInt(1)
		for ; num > 1; num-- {
			res.Mul(big.NewInt(int64(num)), res)
		}
		return res
	}
	factorialN := f(n)
	factorialK := big.NewInt(1).Mul(f(k-1), f(n-k+1))
	res := big.NewInt(1)
	res.DivMod(factorialN, factorialK, big.NewInt(10^9+7))
	return int(res.Int64())
}
