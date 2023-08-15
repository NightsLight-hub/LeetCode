/*
Package _699_ModifyGraphEdgeWeights
@Time : 2023/6/9 10:03
@Author : sunxy
@File : solution
@description:
*/
package _699_ModifyGraphEdgeWeights

import "strconv"

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	var edgeLength = make(map[string]int)
	var pm = make(map[int][]int)
	for _, edge := range edges {
		_, ok := pm[edge[0]]
		if !ok {
			pm[edge[0]] = []int{}
		}
		pm[edge[0]] = append(pm[edge[0]], edge[1])
		_, ok = pm[edge[1]]
		if !ok {
			pm[edge[1]] = []int{}
		}
		pm[edge[1]] = append(pm[edge[1]], edge[0])
		key1 := strconv.Itoa(edge[0]) + "," + strconv.Itoa(edge[1])
		key2 := strconv.Itoa(edge[1]) + "," + strconv.Itoa(edge[0])
		edgeLength[key1] = edge[2]
		edgeLength[key2] = edge[2]
	}
	var paths [][]int
	var search func(int, []int)
	search = func(node int, path []int) {
		// pm 表示 key  与  values相连
		// v in pm[node]
		nextNodes, ok := pm[node]
		if !ok {
			return
		}
		for _, nextNode := range nextNodes {
			var loopPath = false
			for _, pnode := range path {
				if pnode == nextNode {
					// 环路
					loopPath = true
					break
				}
			}
			if loopPath {
				continue
			}
			path = append(path, nextNode)
			if nextNode == destination {
				paths = append(paths, path)
			} else {
				search(nextNode, path)
			}
			// pop last
			path = path[:len(path)-1]
		}
	}

	search(source, []int{source})
	if len(paths) == 0 {
		return [][]int{}
	}
	res := -1
	var modifiableEdge int
	var distance int
	fixNum := 0
	for j, path := range paths {
		modifiableEdge = 0
		distance = 0
		for i := 0; i < len(path)-1; i++ {
			key := strconv.Itoa(path[i]) + "," + strconv.Itoa(path[i+1])
			el := edgeLength[key]
			if el == -1 {
				modifiableEdge++
			} else {
				distance += el
			}
		}
		if distance < target && target-distance >= modifiableEdge {
			res = j
			fixNum = target - distance - modifiableEdge
			break
		}
	}
	if res < 0 {
		return [][]int{}
	}
	var ret [][]int
	var flag = true
	for i := 0; i < len(paths[res])-1; i++ {
		key := strconv.Itoa(paths[res][i]) + "," + strconv.Itoa(paths[res][i+1])
		el := edgeLength[key]
		var d int
		if el != -1 {
			d = el
		} else {
			if flag {
				d = 1 + fixNum
				flag = false
			} else {
				d = 1
			}
		}
		ret = append(ret, []int{paths[res][i], paths[res][i+1], d})
	}
	return ret
}
