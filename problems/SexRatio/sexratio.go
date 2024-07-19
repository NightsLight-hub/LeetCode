/*
Package SexRatio
@Time : 2024/7/11 9:20
@Author : sunxy
@File : sexratio
@description:
大灾难。在大灾难后的新世界，世界女王非常关心出生率。
因此，她规定所有家庭都必须有一个女孩，否则将面临巨额罚款。
如果所有的家庭都遵守这个政策——所有家庭在得到一个女孩之前
不断生育，生了女孩之后立即停止生育——那么新一代的性别比例
是多少（假设每次怀孕后生男生女的概率是相等的）？通过逻辑
推理解决这个问题，然后使用计算机进行模拟
*/
package SexRatio

import (
	"fmt"
	"math/rand"
)

const (
	male = iota
	female
)

func born() int {
	if rand.Int()%2 == 0 {
		return male
	} else {
		return female
	}
}

func mockBorn() {
	var femaleCount, maleCount int
	for i := 0; i < 10000; i++ {
		for {
			if born() == female {
				femaleCount++
				fmt.Printf("%d family get girl\n", i)
				println("------------")
				break
			} else {
				maleCount++
				fmt.Printf("%d family get boy\n", i)
			}
		}
	}
	fmt.Printf("feamle count is %d, male count is %d\n", femaleCount, maleCount)
}
