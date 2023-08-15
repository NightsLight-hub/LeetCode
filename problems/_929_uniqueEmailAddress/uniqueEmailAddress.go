/*
Package _929_uniqueEmailAddress
@Time : 2022/6/4 9:16
@Author : sunxy
@File : uniqueEmailAddress
@description:
*/
package _929_uniqueEmailAddress

import "strings"

// memory O(n)   time O(n)
func numUniqueEmails(emails []string) int {
	m := make(map[string]interface{})
	var processEmail func()
	processEmail = func() {
		for _, email := range emails {
			strs := strings.Split(email, "@")
			local, domain := strs[0], strs[1]
			var bs []rune = make([]rune, 0)
			for _, l := range local {
				if l == '.' {
					continue
				} else if l == '+' {
					break
				}
				bs = append(bs, l)
			}
			m[string(bs)+"@"+domain] = nil
		}
	}
	processEmail()
	return len(m)

}
