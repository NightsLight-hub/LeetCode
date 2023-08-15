/*
Package offer2_alienOrder
@Time : 2022/5/31 10:08
@Author : sunxy
@File : alienorder
@description:
*/
package offer2_alienOrder

type node struct {
	Val  byte
	Next *node
}

func (n *node) insert(pre *node, val byte) *node {
	newNode := &node{
		Val:  val,
		Next: pre.Next,
	}
	pre.Next = newNode
	return newNode
}

const (
	notVisited = -1
	visiting   = 0
	visited    = 1
)

var failed bool
var result *node
var startP byte

func alienOrder(words []string) string {
	result = new(node)
	failed = false
	startP = '0'
	// byteMap key  顺序在value 之前,
	byteMap := make(map[byte][]byte)
	// 构建拓扑图
	for i, _ := range words {
		for j := i + 1; j < len(words); j++ {
			cmp(words[i], words[j], byteMap)
		}
	}
	if failed {
		return ""
	}
	// 深度优先
	startP := []byte(words[0])[0]
	visitStatus := make(map[byte]int)
	deepSearch(byteMap, startP, result, visitStatus)
	if failed {
		return ""
	}
	var resStr []byte = make([]byte, 0)
	m := result
	for m.Next != nil {
		resStr = append(resStr, m.Next.Val)
		m = m.Next
	}
	for _, w := range words {
		for _, b := range []byte(w) {
			if _, ok := visitStatus[b]; !ok {
				resStr = append(resStr, b)
			}
		}
	}
	return string(resStr)
}

func deepSearch(bm map[byte][]byte, currentPoint byte, prevNode *node, visitStatus map[byte]int) {
	if failed {
		return
	}
	if v, ok := visitStatus[currentPoint]; ok {
		if v == visiting {
			failed = true
			return
		}
		if v == visited {
			return
		}
	}
	visitStatus[currentPoint] = visiting
	n := result.insert(prevNode, currentPoint)
	for _, p := range bm[currentPoint] {
		deepSearch(bm, p, n, visitStatus)
	}
	visitStatus[currentPoint] = visited
}

func cmp(s1, s2 string, bm map[byte][]byte) {
	if s1 == s2 {
		return
	}
	var l int
	if len(s1) < len(s2) {
		l = len(s1)
	} else {
		l = len(s2)
	}
	mf := false
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			mf = true
			if _, ok := bm[s1[i]]; !ok {
				bm[s1[i]] = make([]byte, 0)
			}
			flag := false
			for _, b := range bm[s1[i]] {
				if b == s2[i] {
					flag = true
					break
				}
			}
			if !flag {
				bm[s1[i]] = append(bm[s1[i]], s2[i])
			}
			if startP == '0' {
				startP = s1[i]
			}
			break
		}
	}
	if !mf && len(s2) < len(s1) {
		failed = true
	}
}
