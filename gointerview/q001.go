/*
Package gointerview
@Time : 2022/8/1 15:09
@Author : sunxy
@File : q001
@description:
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var letter, ch = make(chan bool), make(chan bool)
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func(i int) {
		for range ch {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- true
		}
	}(1)
	go func(i rune) {
		defer wg.Done()
		for {
			if i > 'Z' {
				break
			}
			select {
			case <-letter:
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				ch <- true
			}
		}
	}('A')
	ch <- true
	wg.Wait()
}
