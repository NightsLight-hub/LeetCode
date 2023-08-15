/*
@Time : 2021/5/28 10:39
@Author : sunxy
@File : total_Hamming_Distance
@description:
*/
package total_Hamming_Distance

var cache = make(map[int]int)

// 10^9 < 2^30
func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		ans += c * (n - c)
	}
	return
}
