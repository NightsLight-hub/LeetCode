/*
Package main
@Time : 2022/8/1 16:16
@Author : sunxy
@File : q003
@description:
*/
package main

import (
	"fmt"
)

func main() {
	s := "123467890"
	i, j := 0, len(s)-1
	bs := []rune(s)
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	fmt.Println(string(bs))
}
