/*
Package main
@Time : 2022/8/1 16:26
@Author : sunxy
@File : q004
@description:
问题描述

给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s1 := "abcddef"
	s2 := "dbdafed"
	m := make(map[rune]int)
	for _, c := range s1 {
		if _, ok := m[c]; !ok {
			m[c] = 0
		}
		m[c]++
	}
	for _, c := range s2 {
		if _, ok := m[c]; !ok {
			fmt.Println("false")
			os.Exit(1)
		}
		if strings.Count(s2, string(c)) != m[c] {
			fmt.Println("false")
			os.Exit(1)
		}
	}
	fmt.Println("true")
	os.Exit(0)
}
