/*
Package _732_my_calendar
@Time : 2022/6/6 15:50
@Author : sunxy
@File : solution_test.go
@description:
*/
package _732_my_calendar

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	myCalendarThree := Constructor()
	a := myCalendarThree.Book(10, 20) // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
	a = myCalendarThree.Book(50, 60)  // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
	a = myCalendarThree.Book(10, 40)  // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k 次预订是 2 次预订。
	a = myCalendarThree.Book(5, 15)   // 返回 3 ，剩下的日程安排的最大 k 次预订是 3 次预订。
	a = myCalendarThree.Book(5, 10)   // 返回 3
	a = myCalendarThree.Book(25, 55)  // 返回 3
	fmt.Println(a)
}

// [[],[24,40],[43,50],[27,43],[5,21],[30,40],[14,29],[3,19],[3,14],[25,39],[6,19]]
func TestConstructor2(t *testing.T) {
	myCalendarThree := Constructor()
	a := myCalendarThree.Book(24, 40) // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
	a = myCalendarThree.Book(43, 50)  // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
	a = myCalendarThree.Book(27, 43)  // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k 次预订是 2 次预订。
	a = myCalendarThree.Book(5, 21)   // 返回 2 ，剩下的日程安排的最大 k 次预订是 3 次预订。
	a = myCalendarThree.Book(30, 40)  // 返回 3
	a = myCalendarThree.Book(14, 29)  // 返回 4
	fmt.Println(a)
}
