/*
Package sorted_circular_linked_node
@Time : 2022/6/18 15:07
@Author : sunxy
@File : solution_test.go
@description:
*/
package sorted_circular_linked_node

import (
	"reflect"
	"testing"
)

func buildLink(a []int) *Node {
	head := &Node{
		Val:  a[0],
		Next: nil,
	}
	n := head
	for i := 1; i < len(a); i++ {
		n.Next = &Node{
			Val:  a[i],
			Next: nil,
		}
		n = n.Next
		if i == len(a)-1 {
			n.Next = head
		}
	}
	return head
}

func Test_insert(t *testing.T) {
	type args struct {
		aNode *Node
		x     int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				aNode: buildLink([]int{3, 4, 1}),
				x:     2,
			},
			want: buildLink([]int{3, 4, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.aNode, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
