/*
Package _817_linkedListCom
@Time : 2022/10/12 10:37
@Author : sunxy
@File : solution_test.go
@description:
*/
package _817_linkedListCom

import "testing"

func Test_numComponents(t *testing.T) {
	type args struct {
		head *ListNode
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
				nums: []int{0, 1, 3},
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  3,
									Next: nil,
								},
							},
						},
					},
				},
				nums: []int{3, 4, 0, 2, 1},
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
				nums: []int{1, 0},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numComponents(tt.args.head, tt.args.nums); got != tt.want {
				t.Errorf("numComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
