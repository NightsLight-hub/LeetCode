/*
@Time : 2021/7/3 10:10
@Author : sunxy
@File : solution
@description:
https://leetcode-cn.com/problems/sort-characters-by-frequency/
*/
package _51_Sort_Characters_By_Frequency

import "sort"

type charactor struct {
	c byte
	f int
}
type characotrs []charactor

func (c characotrs) Len() int {
	return len(c)
}

func (c characotrs) Less(i, j int) bool {
	return c[i].f > c[j].f
}

func (c characotrs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func frequencySort(s string) string {
	res := ""
	bsf := make(map[byte]int)
	bs := []byte(s)
	for _, b := range bs {
		bsf[b]++
	}
	var cs characotrs
	for k, v := range bsf {
		cs = append(cs, charactor{
			c: k,
			f: v,
		})
	}
	sort.Sort(cs)
	for _, c := range cs {
		for i := 0; i < c.f; i++ {
			res = res + string(c.c)
		}
	}
	return res
}
