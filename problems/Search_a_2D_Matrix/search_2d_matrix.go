/*
@Time : 2021/5/27 16:04
@Author : sunxy
@File : search_2d_matrix
@description:
*/
package Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, 0)
	for _, m1 := range matrix {
		for _, i := range m1 {
			arr = append(arr, i)
		}
	}
	return BinarySearch(&arr, 0, len(arr)-1, target)
}

func BinarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int) bool {
	if leftIndex > rightIndex {
		return false
	}

	//先找到中间下标
	midddle := (leftIndex + rightIndex) / 2
	if (*arr)[midddle] > findVal {
		//说明要查找的数在左边  就应该向 leftIndex ---- (middle - 1)再次查找
		return BinarySearch(arr, leftIndex, midddle-1, findVal)
	} else if (*arr)[midddle] < findVal {
		//如果 arr[middle] < findVal , 就应该向 middel+1---- rightIndex
		return BinarySearch(arr, midddle+1, rightIndex, findVal)
	} else {
		return true
	}
}
