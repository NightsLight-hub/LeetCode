/*
@Time : 2021/7/25 7:32
@Author : sunxy
@File : solution
@description:
*/
package __1743_Restore_the_Array_From_Adjacent_Pairs

func restoreArray(adjacentPairs [][]int) []int {
	var adjacentNumsMap = make(map[int][]int)
	var answer []int
	for _, pair := range adjacentPairs {
		adjacentNumsMap[pair[0]] = append(adjacentNumsMap[pair[0]], pair[1])
		adjacentNumsMap[pair[1]] = append(adjacentNumsMap[pair[1]], pair[0])
	}

	for k, v := range adjacentNumsMap {
		if len(v) == 1 {
			answer = append(answer, k)
			answer = append(answer, v[0])
			break
		}
	}
	for len(answer) < len(adjacentPairs)+1 {
		n := answer[len(answer)-1]
		nL := answer[len(answer)-2]
		if nL == adjacentNumsMap[n][0] {
			answer = append(answer, adjacentNumsMap[n][1])
		} else {
			answer = append(answer, adjacentNumsMap[n][0])
		}
	}
	return answer

}

//
//func restoreArray(adjacentPairs [][]int) []int {
//	l := list.New()
//	ans := list.New()
//	ans.PushFront(adjacentPairs[0][0])
//	ans.PushBack(adjacentPairs[0][1])
//	//usedIndex = append(usedIndex, 0)
//	for i := 1; i < len(adjacentPairs); i++ {
//		pair := adjacentPairs[i]
//		l.PushBack(pair)
//	}
//	for l.Len() != 0 {
//		node := l.Front()
//		for node != nil {
//			pair := node.Value.([]int)
//			num1, num2 := pair[0], pair[1]
//			flag := false
//			if num1 == ans.Front().Value.(int) {
//				ans.PushFront(num2)
//				flag = true
//			} else if num2 == ans.Front().Value.(int) {
//				ans.PushFront(num1)
//				flag = true
//			} else if num1 == ans.Back().Value.(int) {
//				ans.PushBack(num2)
//				flag = true
//			} else if num2 == ans.Back().Value.(int) {
//				ans.PushBack(num1)
//				flag = true
//			}
//			tmpNode := node
//			node = node.Next()
//			if flag{
//				l.Remove(tmpNode)
//			}
//		}
//	}
//	node := ans.Front()
//	var answer []int
//	for node !=nil{
//		answer = append(answer, node.Value.(int))
//		node = node.Next()
//	}
//	return answer
//}
