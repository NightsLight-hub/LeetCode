/*
Package _331_isValidSerialization
@Time : 2024/3/31 8:45
@Author : sunxy
@File : isValidSerialization
@description: https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/
*/

package _331_isValidSerialization

func isValidSerialization(preorder string) bool {
	slots := 1
	for i := 0; i < len(preorder); i++ {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			continue
		}
		if preorder[i] == '#' {
			slots--
		} else {
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
			slots++
		}
	}
	return slots == 0
}
