/*
Package _1656_orderedStream
@Time : 2022/8/16 9:21
@Author : sunxy
@File : solution
@description:
*/
package _1656_orderedStream

// type element struct {
//     id    int
//     value string
// }

type OrderedStream struct {
	stream []string
	ptr    int
	size   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n+1),
		ptr:    1,
		size:   n + 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	result := make([]string, 0)
	this.stream[idKey] = value
	if idKey == this.ptr {
		for i := this.ptr; i < this.size && this.stream[i] != ""; i++ {
			result = append(result, this.stream[i])
			this.ptr++
		}
	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
