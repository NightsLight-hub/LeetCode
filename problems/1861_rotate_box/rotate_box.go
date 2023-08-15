/*
@Time : 2021/6/2 17:03
@Author : sunxy
@File : rotate_box
@description:
*/
package _1861_rotate_box

import (
	"fmt"
)

func rotateTheBox(box [][]byte) [][]byte {
	stone := []byte("#")[0]
	blank := []byte(".")[0]
	//wall := []byte("*")[0]
	//print(box)
	newBox := rotateMatrix(box)

	height := len(newBox)
	width := len(newBox[0])
	for j := 0; j < width; j++ {
		q := []int{}
		for i := height - 1; i >= 0; i-- {
			if newBox[i][j] == blank {
				q = append(q, i)
			} else if newBox[i][j] == stone && len(q) != 0 {
				newBox[q[0]][j] = stone
				newBox[i][j] = blank
				q = q[1:]
				q = append(q, i)
			} else {
				q = []int{}
			}

		}
	}
	//print(newBox)
	return newBox
}

func rotateMatrix(box [][]byte) [][]byte {
	height := len(box)
	width := len(box[0])
	newBox := make([][]byte, width)
	for i, _ := range newBox {
		newBox[i] = make([]byte, height)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newBox[j][height-i-1] = box[i][j]
		}
	}
	return newBox
}

func print(box [][]byte) {
	fmt.Println()
	for i := 0; i < len(box); i++ {
		fmt.Printf("%s\n", box[i])
	}
}
