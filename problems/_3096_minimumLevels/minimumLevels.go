package _3096_minimumLevels

// https://leetcode.cn/problems/minimum-levels-to-gain-more-points/
// 官方解更好

// O(N)
func minimumLevels(possible []int) int {
	// scoreA = p[0],
	// for i in [1, len-1]   scoreB
	scoreA := 0
	if possible[0] == 0 {
		scoreA--
	} else {
		scoreA++
	}
	scoreB := 0
	for _, p := range possible[1:] {
		if p == 0 {
			scoreB--
		} else {
			scoreB++
		}
	}

	// scoreA > scoreB   return
	if scoreA > scoreB {
		return 1
	}

	// for i in [1, len-1]   scoreB
	// if p[i] == 0  scoreA++, scoreB--
	// else scoreA--, scroeB++
	// compare
	for i, p := range possible[1:] {
		if p == 0 {
			scoreB++
			scoreA--
		} else {
			scoreB--
			scoreA++
		}
		if scoreA > scoreB {
			// 不能让B 一关没有
			if i+2 == len(possible) {
				return -1
			} else {
				return i + 2
			}
		}
	}
	return -1
}
