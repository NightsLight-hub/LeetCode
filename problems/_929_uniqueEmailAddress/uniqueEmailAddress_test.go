/*
Package _929_uniqueEmailAddress
@Time : 2022/6/4 9:16
@Author : sunxy
@File : uniqueEmailAddress_test.go
@description:
*/
package _929_uniqueEmailAddress

import "testing"

func Test_numUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}},
			want: 2,
		},
		{
			name: "t2",
			args: args{emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
