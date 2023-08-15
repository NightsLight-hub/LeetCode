/*
@Time : 2021/6/2 17:12
@Author : sunxy
@File : rotate_box_test.go
@description:
*/
package _1861_rotate_box

import (
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	str := "##*.*."
	str1 := "###*.."
	str2 := "###.#."
	box := [][]byte{[]byte(str), []byte(str1), []byte(str2)}
	rotateTheBox(box)
}
