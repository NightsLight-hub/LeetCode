package main

import (
	"fmt"
	"sync"
)

func writeChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func readChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		num, ok := <-ch
		if !ok {
			break
		}
		square := num * num
		fmt.Println(square)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go writeChannel(ch, &wg)
	go readChannel(ch, &wg)

	wg.Wait()
}

var ch = make(chan int, 1)

func a() {
	for i := 1; i <= 100; i++ {
		ch <- i
	}

}

func b() {
	for value := range ch {
		square := value * value
		fmt.Println(square)
	}
}

func c() {
	go a()
	go b()
	select {}
}
