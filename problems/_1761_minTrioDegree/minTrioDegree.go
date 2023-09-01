/*
Package _1761_minTrioDegree
@Time : 2023/8/31 9:20
@Author : sunxy
@File : minTrioDegree
@description: https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph/
实现不正确，应该用官方解
用矩阵存储，如果 vec[i][j]==1 && vec[j][k] == 1 && vec[i][k] == 1 三元组
为了减少重复计算，让 i < j < k
*/
package _1761_minTrioDegree

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// walk 不能保证找到三角形， 有可能图非常复杂，有非常多的分支，也许可以用广度优先？保存三步
func minTrioDegree(n int, edges [][]int) int {
	// make a map, map's key is the point and the value is the array of points that it can reach
	res := math.MaxInt32
	var m = make(map[int][]int)
	var points = make(map[int]struct{})
	for i := range edges {
		p1, p2 := edges[i][0], edges[i][1]
		if _, ok := m[p1]; !ok {
			m[p1] = make([]int, 0)
		}
		m[p1] = append(m[p1], p2)
		if _, ok := m[p2]; !ok {
			m[p2] = make([]int, 0)
		}
		m[p2] = append(m[p2], p1)
		points[p1] = struct{}{}
		points[p2] = struct{}{}
	}
	tripleMetas := make(map[string][]int)
	// cl2 是c 前两步的值， cl1 是前一步的值， c是当前值
	cl2, cl1, c := -1, -1, edges[0][0]
	var walk func(cur, cl1, cl2 int)
	// 走过的路
	var path = make(map[int]struct{})
	path[edges[0][0]] = struct{}{}
	walk = func(cur, cl1, cl2 int) {
		curNexts, ok := m[cur]
		if !ok {
			return
		}
		// 当前cur已经走过
		path[cur] = struct{}{}
		delete(points, cur)
		for _, curNext := range curNexts {
			// 下一个可走的点 与cl2相同，说明走出一个三角形
			if curNext == cl2 {
				tripleMeta := []int{cur, cl1, cl2}
				sort.Ints(tripleMeta)
				key := ""
				for i := range tripleMeta {
					key += strconv.Itoa(tripleMeta[i]) + "-"
				}
				if _, ok := tripleMetas[key]; !ok {
					du := 0
					tripleMetas[key] = tripleMeta
					// 计算度
					du += len(m[tripleMeta[0]])
					du += len(m[tripleMeta[1]])
					du += len(m[tripleMeta[2]])
					du -= 6
					if du < res {
						res = du
					}
				}

			}
			if _, ok := path[curNext]; !ok {
				// 没走过的路
				walk(curNext, cur, cl1)
			}
		}
	}
	walk(c, cl1, cl2)
	for len(points) != 0 {
		for k := range points {
			walk(k, -1, -1)
			break
		}
	}
	fmt.Printf("%#v\n", tripleMetas)
	if res == math.MaxInt32 {
		res = -1
	}
	return res
}
