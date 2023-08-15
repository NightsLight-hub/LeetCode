/*
Package _611_miceAndCheese
@Time : 2023/6/7 10:02
@Author : sunxy
@File : Solution_test.go
@description:
*/
package _611_miceAndCheese

import "testing"

func Test_miceAndCheese(t *testing.T) {
	if 17 != miceAndCheese([]int{1, 4, 2, 1, 1}, []int{5, 6, 1, 2, 5}, 3) {
		t.Fatalf("1 failed")
	}
	if 21 != miceAndCheese([]int{4, 1, 5, 3, 3}, []int{3, 4, 4, 5, 2}, 3) {
		t.Fatalf("1 failed")
	}
	if 15 != miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2) {
		t.Fatalf("1 failed")
	}
	if 2 != miceAndCheese([]int{1, 1}, []int{1, 1}, 2) {
		t.Fatalf("2 failed")
	}
	if 4 != miceAndCheese([]int{2, 1}, []int{1, 2}, 1) {
		t.Fatalf("3 failed")
	}
	if 18 != miceAndCheese([]int{10, 9}, []int{9, 1}, 1) {
		t.Fatalf("4 failed")
	}
	if 10 != miceAndCheese([]int{3, 1, 1, 3}, []int{3, 2, 3, 1}, 3) {
		t.Fatalf("5 failed")
	}
}
