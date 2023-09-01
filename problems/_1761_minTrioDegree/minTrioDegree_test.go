/*
Package _1761_minTrioDegree
@Time : 2023/8/31 9:50
@Author : sunxy
@File : minTrioDegree_test.go
@description:
*/
package _1761_minTrioDegree

import (
	"fmt"
	"testing"
)

func Test_minTrioDegree(t *testing.T) {
	var res int
	// res := minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}})
	// fmt.Printf("du is %d\n", res)
	// res = minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}})
	// fmt.Printf("du is %d\n", res)
	res = minTrioDegree(17, [][]int{{12, 10}, {12, 16}, {4, 9}, {4, 6}, {14, 1}, {9, 2},
		{17, 6}, {17, 12}, {8, 9}, {11, 14}, {13, 5}, {8, 15}, {13, 11}, {15, 11}, {15, 14},
		{6, 8}, {12, 15}, {14, 12}, {9, 1}, {9, 10}, {10, 5}, {1, 11}, {2, 10}, {15, 1}, {7, 9},
		{14, 2}, {4, 1}, {17, 7}, {3, 17}, {8, 1}, {17, 13}, {10, 13}, {8, 13}, {1, 7}, {2, 6},
		{13, 6}, {7, 2}, {1, 16}, {6, 3}, {6, 9}, {16, 17}, {7, 14}})
	fmt.Printf("du is %d\n", res)
}
