/*
@Time : 2021/5/31 18:10
@Author : sunxy
@File : power_of_four_test.go
@description:
*/
package Power_of_Four

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	t.Logf("%v", isPowerOfFour(4))
	t.Logf("%v", isPowerOfFour(1))
	t.Logf("%v", isPowerOfFour(16))
	t.Logf("%v", isPowerOfFour(1025))
}
