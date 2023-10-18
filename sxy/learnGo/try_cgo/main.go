/*
Package try_cgo
@Time : 2023/9/22 16:08
@Author : sunxy
@File : main
@description:
*/
package main

import "C"
import "fmt"

type Person struct {
	Name string
	Age  int
}

var personInstance = new(Person)

func GetPerson() *Person {
	return personInstance
}

//export SetName
func SetName(name string) {
	GetPerson().Name = name
}

//export SetAge
func SetAge(age int) {
	GetPerson().Age = age
}

//export personToString
func personToString() string {
	return fmt.Sprintf("tdata name %s age %d", GetPerson().Name, GetPerson().Age)
}

func main() {}
