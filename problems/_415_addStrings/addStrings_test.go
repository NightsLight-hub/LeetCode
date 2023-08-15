/*
Package _415_addStrings
@Time : 2023/7/17 10:23
@Author : sunxy
@File : addStrings_test.go
@description:
*/
package _415_addStrings

import (
	"fmt"
	"testing"
)

func Test_addStrings(t *testing.T) {
	fmt.Printf("%s\n", addStrings("123", "456"))
	fmt.Printf("%s\n", addStrings("123", "1111"))
	fmt.Printf("%s\n", addStrings("123", "1111"))
	fmt.Printf("%s\n", addStrings("11", "123"))
	fmt.Printf("%s\n", addStrings("456", "77"))
	fmt.Printf("%s\n", addStrings("0", "0"))
}
