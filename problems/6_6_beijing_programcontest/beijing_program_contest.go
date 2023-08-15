/*
@Time : 2021/6/4 18:00
@Author : sunxy
@File : beijing_program_contest
@description:
*/
package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

var year int
var month int
var monthCount = 12
var monthStartDay = "mon"
var outputs []string

func perpetualCalendar(year int, month int) (int, int) {
	var monthDays int = 0
	var totalDyas int = 0
	var isLeapYear bool = false
	//fmt.Println("query date is:", year, month)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		isLeapYear = true
		//fmt.Println(year, "is leap year")
	} else {
		isLeapYear = false
		//fmt.Println(year, "is not lear year")
	}

	//计算距离1900年的天数
	for i := 1900; i < year; i++ {
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			totalDyas += 366
		} else {
			totalDyas += 365
		}
	}

	//计算截至到当月1号前的总天数
	for j := 1; j <= month; j++ {
		switch j {
		case 1, 3, 5, 7, 8, 10, 12:
			monthDays = 31
			break
		case 2:
			if isLeapYear {
				monthDays = 29
			} else {
				monthDays = 28
			}
			break
		case 4, 6, 9, 11:
			monthDays = 30
			break
		default:
			fmt.Println("input month is error.")
		}

		if j != month {
			totalDyas += monthDays
		}
	}

	//计算当月1号是几
	var weekDay int = 0
	weekDay = 1 + totalDyas%7
	if weekDay == 7 {
		weekDay = 0
	}

	//fmt.Println("weekDay is: ", weekDay)
	return weekDay, monthDays
}

func showMonthCalender(currentOutputs []string, weekDay, monthDays int) {
	magic1 := 0
	magic2 := 0
	if monthStartDay == "mon" {
		magic1 = 37
		magic2 = 1
		currentOutputs[1] = fmt.Sprintf("%s一\t二\t三\t四\t五\t六\t日\t\t", currentOutputs[1])
		// 如果 周一开头，而1号恰好是周日，这里需要修正一下， 0 就是 7
		if weekDay == 0 {
			weekDay = 7
		}
	}

	if monthStartDay == "sun" {
		magic1 = 36
		magic2 = 0
		currentOutputs[1] = fmt.Sprintf("%s日\t一\t二\t三\t四\t五\t六\t\t", currentOutputs[1])
	}
	count := 0
	outputIndex := 2
	lastInputIndex := 0
	for i := magic1 - weekDay; i <= monthDays; i++ {
		currentOutputs[outputIndex] = fmt.Sprintf("%s%d\t", currentOutputs[outputIndex], i)
		count++
	}
	for k := 0; k < weekDay-count-magic2; k++ {
		currentOutputs[outputIndex] = fmt.Sprintf("%s\t", currentOutputs[outputIndex])
	}

	for m := 1; m <= monthDays-count; m++ {
		lastInputIndex++
		currentOutputs[outputIndex] = fmt.Sprintf("%s%d\t", currentOutputs[outputIndex], m)
		//fmt.Printf("%d\t", m)
		if (weekDay+m-magic2)%7 == 0 {
			currentOutputs[outputIndex] = currentOutputs[outputIndex] + "\t"
			lastInputIndex = 0
			outputIndex++
		}
	}
	if lastInputIndex > 0 {
		currentOutputs[outputIndex] = currentOutputs[outputIndex] + "\t"
	}
	for m := lastInputIndex; m > 0 && m < 7; m++ {
		currentOutputs[outputIndex] = currentOutputs[outputIndex] + "\t"
	}
}

func showCalendar(wms [][]int) {
	for i, wm := range wms {
		currentOutPuts := outputs[i/3*7 : i/3*7+7]
		currentyear := year
		currentMonth := month + i
		if currentMonth > 12 {
			currentMonth = currentMonth - 12
			currentyear++
		}

		if (i+1)%3 == 0 {
			currentOutPuts[0] = fmt.Sprintf("%s%d年%d月", currentOutPuts[0], currentyear, currentMonth)
			showMonthCalender(currentOutPuts, wm[0], wm[1])
		} else {
			currentOutPuts[0] = fmt.Sprintf("%s%d年%d月\t\t\t\t\t\t", currentOutPuts[0], currentyear, currentMonth)
			if runtime.GOOS == "linux" {
				currentOutPuts[0] += "\t"
			}
			showMonthCalender(currentOutPuts, wm[0], wm[1])
		}
	}
	//fmt.Println()
}

func outputCalender() {
	outputs = make([]string, (monthCount/3+1)*7)
	months := make([]int, 0)
	wms := make([][]int, monthCount)

	for i := 0; i < monthCount; i++ {
		wms[i] = make([]int, 2)
	}
	for i := 0; i < monthCount; i++ {
		if month+i > 12 {
			months = append(months, (month+i)%12)
		} else {
			months = append(months, month+i)
		}
	}

	var last = months[0]
	for i, m := range months {
		theYear := year
		if m < last {
			theYear++
		}
		wms[i][0], wms[i][1] = perpetualCalendar(theYear, m)
	}

	showCalendar(wms)
	for i, s := range outputs {
		if s != "" {
			fmt.Println(s)
		}
		if (i+1)%7 == 0 {
			fmt.Println()
		}
	}
}

func main() {
	//开启main 函数的解析功能
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("help message? hehe")
		os.Exit(0)
	}
	data, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("help message? hehe")
		os.Exit(0)
	}
	month = data % 100
	year = (data - month) / 100

	if len(os.Args) >= 3 {
		monthCount, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("help message? hehe")
			os.Exit(1)
		}
	}

	if len(os.Args) >= 4 {
		monthStartDay = os.Args[3]
	}

	perpetualCalendar(year, month)
	outputCalender()
	elapsed := time.Since(start)
	fmt.Printf("执行时长：%v, %s", elapsed, "作者：孙星宇")
}
