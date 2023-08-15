/*
Package stack
@Time : 2022/5/28 9:51
@Author : sunxy
@File : stack
@description:
*/
package stack

import (
	"fmt"
	"strings"
)

type Element interface{}

type ElementShow func(element Element) string

type Stack struct {
	content []Element
}

func NewStack() *Stack {
	s := new(Stack)
	s.content = make([]Element, 0)
	return s
}

func (s *Stack) Push(e interface{}) *Stack {
	s.content = append(s.content, e)
	return s
}

func (s *Stack) Pop() (ret Element) {
	if s.Len() == 0 {
		return nil
	}
	ret = s.content[s.Len()-1]
	s.content = s.content[:len(s.content)-1]
	return
}

func (s *Stack) Head() Element {
	if s.Len() == 0 {
		return nil
	}
	return s.content[0]
}

func (s *Stack) Len() int {
	return len(s.content)
}

func (s *Stack) IsEmpty() interface{} {
	return s.Len() == 0
}

func (s *Stack) Show(se ElementShow) string {
	ret := ""
	for _, e := range s.content {
		ret = ret + " " + se(e)
	}
	ret = strings.TrimSpace(ret)
	fmt.Println(ret)
	return ret
}
