/*
Package _636_ExclusiveTimeofFunctions
@Time : 2022/8/7 9:33
@Author : sunxy
@File : solution
@description:
*/
package _636_ExclusiveTimeofFunctions

import (
	"strconv"
	"strings"
)

const (
	END   = "end"
	START = "start"
)

type entity struct {
	name  int
	time  int
	state string
}

type stack struct {
	values []*entity
}

func newStack() *stack {
	s := new(stack)
	s.values = make([]*entity, 0)
	return s
}

func (s *stack) push(e *entity) {
	s.values = append(s.values, e)
}

func (s *stack) pop() (e *entity) {
	e = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return
}
func (s *stack) head() (e *entity) {
	if len(s.values) != 0 {
		return s.values[len(s.values)-1]
	}
	return nil
}
func (s *stack) len() int {
	return len(s.values)
}
func stringToInt(s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return a
}

func record(result map[int]int, name int, duration int, maxName *int) {
	_, ok := result[name]
	if ok {
		result[name] = result[name] + duration
	} else {
		result[name] = duration
	}
	if *maxName < name {
		*maxName = name
	}
}

func exclusiveTime(n int, logs []string) []int {
	var maxName int = 0
	// 已经被记录的其他函数执行时间，统计某个函数独占时间，需要用结尾 +1 - 开头 - 已经被记录过的时长
	result := make(map[int]int)
	s := newStack()
	t := 0
	for _, log := range logs {
		tmp := strings.Split(log, ":")
		e := &entity{
			name:  stringToInt(tmp[0]),
			time:  stringToInt(tmp[2]),
			state: tmp[1],
		}
		// s.head 指向当前正在执行的函数名
		if e.state == START {
			if s.head() != nil {
				record(result, s.head().name, e.time-t, &maxName)
				t = e.time
			}
			s.push(e)
		} else {
			s.pop()
			record(result, e.name, e.time-t+1, &maxName)
			t = e.time + 1
		}
	}

	fr := make([]int, maxName+1)
	for k, v := range result {
		fr[k] = v
	}
	return fr
}
