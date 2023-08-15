/*
@Time : 2021/7/24 7:36
@Author : sunxy
@File : solution
@description:
*/
package __1736_Latest_Time_by_Replacing_Hidden_Digits

var num0 = []byte("0")[0]
var num1 = num0 + 1
var num2 = num0 + 2

func maximumTime(time string) string {
	bs := []byte(time)
	for i, _ := range bs {
		if bs[i] == []byte("?")[0] {
			if i == 0 {
				if bs[1] >= []byte("4")[0] && bs[1] <= []byte("9")[0] {
					bs[i] = []byte("1")[0]
				} else {
					bs[i] = []byte("2")[0]
				}
			}
			if i == 1 {
				switch bs[0] {
				case num0, num1:
					bs[1] = []byte("9")[0]
				case num2:
					bs[1] = []byte("3")[0]
				}
			}
			if i == 3 {
				bs[3] = []byte("5")[0]
			}
			if i == 4 {
				bs[4] = []byte("9")[0]
			}
		}
	}
	return string(bs)
}
