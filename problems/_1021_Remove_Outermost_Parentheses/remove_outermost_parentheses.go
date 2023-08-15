/*
Package _1021_Remove_Outermost_Parentheses
@Time : 2022/5/28 9:46
@Author : sunxy
@File : remove_outermost_parentheses
@description:
*/
package _1021_Remove_Outermost_Parentheses

import (
	"fmt"
	"strings"
)

type Element interface{}

type Stack struct {
	content []Element
}

func newStack() *Stack {
	s := new(Stack)
	s.content = make([]Element, 0)
	return s
}

func (s *Stack) push(e interface{}) *Stack {
	s.content = append(s.content, e)
	return s
}

func (s *Stack) pop() (ret Element) {
	if s.len() == 0 {
		return nil
	}
	ret = s.content[s.len()-1]
	s.content = s.content[:len(s.content)-1]
	return
}

func (s *Stack) head() Element {
	if s.len() == 0 {
		return nil
	}
	return s.content[0]
}

func (s *Stack) len() int {
	return len(s.content)
}

func (s *Stack) isEmpty() interface{} {
	return s.len() == 0
}

func (s *Stack) show(se func(element Element) string) string {
	ret := ""
	for _, e := range s.content {
		ret = ret + " " + se(e)
	}
	ret = strings.TrimSpace(ret)
	fmt.Println(ret)
	return ret
}

func removeOuterParentheses(s string) (ret string) {
	if s == "" {
		return s
	}
	st := newStack()
	bts := []byte(s)
	leftIndex := 0
	for i, b := range bts {
		if b == '(' {
			st.push(b)
		} else {
			st.pop()
			if st.len() == 0 {
				originStr := string(bts[leftIndex : i+1])
				bts := []byte(originStr)
				ret += string(bts[1 : len(bts)-1])
				leftIndex = i + 1
			}
		}
	}
	return ret
}
