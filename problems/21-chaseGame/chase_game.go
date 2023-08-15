/*
@Time : 2021/6/2 11:01
@Author : sunxy
@File : chase_game
@description:
*/
package _21_chaseGame

// The distance from startPoint to each point
type StartPointMap struct {
	startPoint int
	distance   map[int]int
}

var graph = make(map[int][]int)

func initStartPointMaps(edges [][]int) map[int]StartPointMap {
	startPointMaps := make(map[int]StartPointMap)
	for _, v := range edges {
		spm := StartPointMap{
			startPoint: v[0],
		}
	}
}

func bfs(start, end int) int {
	var q = make([]int, 0)

}

func initGraph(edges [][]int) {
	for _, v := range edges {
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = make([]int, 0)
		}
		graph[v[0]] = append(graph[v[0]], v[1])
		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = make([]int, 0)
		}
		graph[v[1]] = append(graph[v[1]], v[0])
	}
}
func chaseGame(edges [][]int, startA int, startB int) int {
	initGraph(edges)

}
