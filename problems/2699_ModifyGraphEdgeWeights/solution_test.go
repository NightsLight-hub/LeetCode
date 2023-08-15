/*
Package _699_ModifyGraphEdgeWeights
@Time : 2023/6/9 11:00
@Author : sunxy
@File : solution_test.go
@description:
*/
package _699_ModifyGraphEdgeWeights

import (
	"fmt"
	"testing"
)

func Test_modifiedGraphEdges(t *testing.T) {
	// n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0, destination = 1, target = 5
	// var edges [][]int
	// edges = append(edges, []int{4, 1, -1})
	// edges = append(edges, []int{2, 0, -1})
	// edges = append(edges, []int{0, 3, -1})
	// edges = append(edges, []int{4, 3, -1})
	// fmt.Printf("%#v", modifiedGraphEdges(5, edges, 0, 1, 5))

	var edges [][]int
	// edges = append(edges, []int{0, 1, -1})
	// edges = append(edges, []int{1, 2, -1})
	// edges = append(edges, []int{1, 3, -1})
	// edges = append(edges, []int{2, 3, -1})
	// edges = append(edges, []int{3, 4, -1})
	// fmt.Printf("%#v", modifiedGraphEdges(5, edges, 0, 4, 10))

	edges = [][]int{}
	edges = append(edges, []int{1, 0, 4})
	edges = append(edges, []int{1, 2, 3})
	edges = append(edges, []int{2, 3, 5})
	edges = append(edges, []int{0, 3, -1})
	fmt.Printf("%#v", modifiedGraphEdges(4, edges, 0, 2, 6))

}
