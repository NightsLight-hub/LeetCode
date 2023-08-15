/*
@Time : 2021/5/31 18:07
@Author : sunxy
@File : power_of_four
@description:
*/
package Power_of_Four

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	r := n % 10
	if r != 4 && r != 6 {
		return false
	}
	r = 1
	for r < n {
		r = r * 4
	}
	if r != n {
		return false
	}
	return true
}
