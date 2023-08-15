/*
Package _833_findReplaceString
@Time : 2023/8/15 9:06
@Author : sunxy
@File : findReplaceString
@description:
*/
package _833_findReplaceString

import "strings"

// 以abcd 为例，
// abcd 0，2 替换为 eee 和 ffff
/**
则si 为 [0,-1,1,1]
ss ["eee", "ffff"]

扫描si 如果 si[i] == -1 ，那说明s 的i 没有被替换，用s[i]
如果 s[i] != -1, s[i]替换过， 替换的值 在ss[si[i]]
平均复杂度
indices的长度为n， sources中的字符串的平均长度m，则时间复杂度 O(n*m)
空间复杂度 O(n + N*M)
*/

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// si 与 s长度一致，用来记录对s 下表操作后，替换的值的index， 值存放在ss中
	si := make([]int, len(s))
	for i := range si {
		si[i] = -1
	}
	// ss存放替换的值， ss的索引 是si中存放的值
	var ss []string
	for i, v := range indices {
		source, target := sources[i], targets[i]
		if strings.HasPrefix(s[v:], source) {
			ss = append(ss, target)
			for k := 0; k < len(source); k++ {
				si[v+k] = len(ss) - 1
			}
		}
	}
	var sb strings.Builder
	for i := 0; i < len(si); {
		if si[i] == -1 {
			sb.WriteByte([]byte(s)[i])
			i++
		} else {
			temp := ss[si[i]]
			j := i + 1
			for ; j < len(si); j++ {
				if si[j] != si[i] {
					break
				}
			}
			i = j
			sb.WriteString(temp)
		}
	}
	return sb.String()
}
