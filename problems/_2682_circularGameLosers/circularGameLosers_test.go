/*
Package _2682_circularGameLosers
@Time : 2023/8/16 9:48
@Author : sunxy
@File : circularGameLosers_test.go
@description:
*/
package _2682_circularGameLosers

import (
	"reflect"
	"testing"
)

func Test_circularGameLosers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{
				n: 5,
				k: 2,
			},
			want: []int{4, 5},
		},
		{
			name: "t2",
			args: args{
				n: 4,
				k: 4,
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularGameLosers(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularGameLosers() = %v, want %v", got, tt.want)
			}
		})
	}
}
