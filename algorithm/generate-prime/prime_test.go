/*
Package generate_prime
@Time : 2024/7/10 16:22
@Author : sunxy
@File : prime_test.go
@description:
*/
package generate_prime

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	generate(100)
}

var results []string

func generatePermutations(data []rune, prefix string, remaining map[rune]int) {
	// 当remaining为空时，我们得到了一个完整的排列
	if len(prefix) == len(data) {
		results = append(results, prefix)
	} else {
		// 遍历remaining的键
		for char, count := range remaining {
			// 如果当前字符的数量大于0，我们才考虑使用它
			if count > 0 {
				// 减少当前字符的数量
				remaining[char]--
				// 递归调用generatePermutations
				generatePermutations(data,
					prefix+string(char), remaining)
				// 回溯：恢复remaining的状态
				remaining[char]++
			}
		}
	}
}

func Test_a(t *testing.T) {
	data := []rune{'a', 'b', 'c'}
	// 初始化remaining，记录每个字符出现的次数
	remaining := make(map[rune]int)
	for _, char := range data {
		remaining[char] += len(data)
	}
	results = []string{}
	// 调用generatePermutations函数，初始前缀为空字符串
	generatePermutations(data, "", remaining)
	// 输出结果
	fmt.Println(results)
}
