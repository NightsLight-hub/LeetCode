/*
Package _struct
@Time : 2023/9/21 13:32
@Author : sunxy
@File : ss
@description:
*/
package _struct

import (
	"fmt"
	"testing"
)

func Test_a1(t *testing.T) {
	a := Animal{name: "aaa"}
	c := Cat{Animal: a}
	c2 := c
	d := Dog{Animal: &a}
	d2 := Dog{Animal: &a}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", d)

	// 修改指针嵌套结构的内容
	d.name = "bbb"

	fmt.Printf("c name  %+v\n", c.name)
	fmt.Printf("d name %+v\n", d.name)
	fmt.Printf("d2 name %+v\n", d2.name)

	// 修改c2值嵌套结构的内容
	c.name = "ccc"
	fmt.Printf("c name  %+v\n", c.name)
	fmt.Printf("c2 name %+v\n", c2.name)

	fmt.Println("在函数内修改值嵌套结构属性")
	modifyAndPrintName(c2)
	fmt.Printf("c2 name %+v\n", c2.name)

}

func modifyAndPrintName(obj Cat) {
	obj.name = "obj"
	fmt.Printf("name  %+v\n", obj.name)
}

type Animal struct {
	name string
}

type Cat struct {
	Animal
}

type Dog struct {
	*Animal
}
