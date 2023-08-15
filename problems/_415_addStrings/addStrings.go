/*
Package _415_addStrings
@Time : 2023/7/17 10:18
@Author : sunxy
@File : addStrings
@description:
*/
package _415_addStrings

import "strconv"

func addStrings(num1 string, num2 string) string {
	a, b := []byte(num1), []byte(num2)
	var ret = ""
	f := 0
	for {
		a1, alen0 := val(&a)
		b1, blen0 := val(&b)
		if alen0 && blen0 && f == 0 {
			break
		}
		m := a1 + b1 + f
		if m > 9 {
			f = 1
			m %= 10
		} else {
			f = 0
		}
		ret = strconv.Itoa(m) + ret
	}
	return ret
}

func val(n *[]byte) (ret int, len0 bool) {
	if len(*n) == 0 {
		return 0, true
	}
	ret = int((*n)[len(*n)-1] - '0')
	*n = (*n)[:len(*n)-1]
	return ret, false
}
