/*
Package _1656_orderedStream
@Time : 2022/8/16 9:30
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1656_orderedStream

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	os := Constructor(5)
	ret := os.Insert(3, "ccccc") // 插入 (3, "ccccc")，返回 []
	fmt.Printf("%+v", ret)
	ret = os.Insert(1, "aaaaa") // 插入 (1, "aaaaa")，返回 ["aaaaa"]
	fmt.Printf("%+v", ret)
	ret = os.Insert(2, "bbbbb") // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"]
	fmt.Printf("%+v", ret)
	ret = os.Insert(5, "eeeee") // 插入 (5, "eeeee")，返回 []
	fmt.Printf("%+v", ret)
	ret = os.Insert(4, "ddddd") // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"]
	fmt.Printf("%+v", ret)
}
