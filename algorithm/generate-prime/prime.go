/*
Package generate_prime
@Time : 2024/7/10 16:10
@Author : sunxy
@File : prime
@description:
*/
package generate_prime

import (
	"fmt"
	"math"
)

var ff []bool

func generate(max int) {
	ff = make([]bool, max+1)
	p := 2
	sqrtResult := int(math.Sqrt(float64(max)))
	for p <= sqrtResult {
		reduce(p, max)
		p = getNextPrime(p)
	}
	for i := 1; i < len(ff); i++ {
		if !ff[i] {
			fmt.Printf("%d\t", i)
		}
	}
}

func reduce(prime, max int) {
	i := prime + 1
	for i <= max {
		if i != prime && i%prime == 0 {
			ff[i] = true
		}
		i++
	}
}

func getNextPrime(prime int) int {
	for i := prime + 1; i < len(ff); i++ {
		if !ff[i] {
			return i
		}
	}
	return math.MaxInt32
}
