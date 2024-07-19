/*
Package generate_prime
@Time : 2024/7/11 16:05
@Author : sunxy
@File : prime2
@description:
*/
package generate_prime

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func startGeneratePrimes() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 100; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
