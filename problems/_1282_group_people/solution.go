/*
Package _1282_group_people
@Time : 2022/8/12 9:08
@Author : sunxy
@File : solution
@description:
*/
package _1282_group_people

type group struct {
	values []int
	length int
	size   int
}

func newGroup(size int) *group {
	return &group{
		values: make([]int, 0),
		length: 0,
		size:   size,
	}
}

func (g *group) add(p int) {
	if g.length >= g.size {
		panic("impossible!")
	}
	g.values = append(g.values, p)
	g.length++
}

var groups []*group

func getGroup(gsize int) *group {
	for _, g := range groups {
		if g.size == gsize && g.length < g.size {
			return g
		}
	}
	return nil
}

func result() [][]int {
	result := make([][]int, 0)
	for _, g := range groups {
		result = append(result, g.values)
	}
	return result
}

func groupThePeople(groupSizes []int) [][]int {
	groups = make([]*group, 0)
	for people, gsize := range groupSizes {
		g := getGroup(gsize)
		if g == nil {
			g = newGroup(gsize)
			groups = append(groups, g)
		}
		g.add(people)
	}
	return result()
}
