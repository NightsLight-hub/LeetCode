/*
@Time : 2021/7/18 7:27
@Author : sunxy
@File : solution
@description:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[

	["ate","eat","tea"],
	["nat","tan"],
	["bat"]

]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package _10_02_Group_Anagrams_LCCI

import "container/list"

func groupAnagrams(strs []string) [][]string {
	strList := list.New()
	for _, v := range strs {
		strList.PushBack(v)
	}
	for strList.Len() != 0 {
		frontElement := strList.Front()
		if frontElement == nil {
			break
		}
		strList.Remove(frontElement)
		targetByteMap := make(map[byte]int)
		for _, b := range []byte(frontElement.Value.(string)) {
			targetByteMap[b]++
		}
		for {
			cursorEle := strList.Front()
			if cursorEle == nil {
				break
			}
			cursorByteMap := make(map[byte]int)
			for _, b := range []byte(cursorEle.Value.(string)) {
				cursorByteMap[b]++
			}
			if isSame(targetByteMap, cursorByteMap) {

			}
		}
	}
}

func isSame(a map[byte]int, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		bv, ok := b[k]
		if !ok {
			return false
		}
		if bv != v {
			return false
		}
	}
	return true
}
