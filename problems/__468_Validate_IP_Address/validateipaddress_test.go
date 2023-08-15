/*
Package __468_Validate_IP_Address
@Time : 2022/5/29 9:32
@Author : sunxy
@File : validateipaddress_test.go
@description:
*/
package __468_Validate_IP_Address

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		queryIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{queryIP: "172.16.254.1"},
			want: "IPv4",
		},
		{
			name: "t2",
			args: args{queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
		{
			name: "t3",
			args: args{queryIP: "256.256.256.256"},
			want: "Neither",
		},
		{
			name: "t4",
			args: args{queryIP: "01.01.01.01"},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.queryIP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
