/*
Package _2591_distMoney
@Time : 2023/9/22 17:18
@Author : sunxy
@File : distMoney
@description:
*/
package _2591_distMoney

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if children*8 == money {
		return children
	}
	n := children - 1
	for n > 0 {
		if 8*n < money {
			if children-n == 1 && money-8*n == 4 {
				// 刚好剩一个人，剩4块钱
				n--
			} else if money-8*n < children-n {
				// 剩下的钱分不给没人一块
				n--
			} else {
				return n
			}
		} else {
			n--
		}
	}
	return n
}
