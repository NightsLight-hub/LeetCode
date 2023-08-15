/*
@Time : 2021/7/26 16:28
@Author : sunxy
@File : client_test.go
@description:
*/
package testCh

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	test()
}

func TestDeployMaster(t *testing.T) {
	result := fmt.Sprintf("cd kubernetes && sed -i 's/.*BASEDIR=.*/BASEDIR=\\%s/' install.cfg", "/afdsfasfd")
	t.Log(result)
}
