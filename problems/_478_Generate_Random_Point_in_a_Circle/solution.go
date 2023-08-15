/*
Package _478_Generate_Random_Point_in_a_Circle
@Time : 2022/6/5 10:10
@Author : sunxy
@File : solution
@description:
*/
package _478_Generate_Random_Point_in_a_Circle

import "math/rand"

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius, xCenter, yCenter float64) Solution {
	return Solution{radius, xCenter, yCenter}
}

func (s *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1 // [-1,1) 的随机数
		if x*x+y*y < 1 {
			return []float64{s.xCenter + x*s.radius, s.yCenter + y*s.radius}
		}
	}
}
