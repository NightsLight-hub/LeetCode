/*
Package _1417_reformat_string
@Time : 2022/8/11 9:24
@Author : sunxy
@File : solution
@description:
*/
package _1417_reformat_string

import "strings"

func reformat(s string) string {
	var isNum = func(s rune) bool {
		return s <= '9' && s >= '0'
	}
	r1 := make([]rune, 0)
	r2 := make([]rune, 0)
	for _, a := range []rune(s) {
		if isNum(a) {
			r1 = append(r1, a)
		} else {
			r2 = append(r2, a)
		}
	}
	if len(r1) != len(r2) && (len(r1)-len(r2) != 1 && len(r1)-len(r2) != -1) {
		return ""
	}
	var sb strings.Builder
	minlen := len(r1)
	t := "num"
	if minlen > len(r2) {
		minlen = len(r2)
		t = "letter"
	}
	for i := 0; i < minlen; i++ {
		sb.WriteRune(r1[i])
		sb.WriteRune(r2[i])
	}
}
