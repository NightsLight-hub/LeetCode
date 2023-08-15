/*
Package _478_Generate_Random_Point_in_a_Circle
@Time : 2022/6/5 10:19
@Author : sunxy
@File : solution_test.go
@description:
*/
package _478_Generate_Random_Point_in_a_Circle

import (
	"math"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor(10, 5, -7.5)
	for i := 0; i < 10; i++ {
		p1 := s.RandPoint()
		if !check(p1, 10, 0, 0) {
			t.Fatalf("%f %f", p1[0], p1[1])
		}
	}
}

func check(point []float64, radis, x, y float64) bool {
	xd := point[0] - x
	yd := point[0] - y
	length := math.Sqrt(xd*xd + yd*yd)
	println(length)
	return length <= radis
}
