/*
@Time : 2021/7/16 10:07
@Author : sunxy
@File : generator
@description:
*/
package learnIO

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateFile() {
	f, err := os.OpenFile("E:\\workspace\\a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1e5; i++ {
		n := rand.Int() % 1000
		f.Write([]byte(strconv.Itoa(n) + " "))
	}
}

func readFile() {
	f, err := os.Open("E:\\workspace\\a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bureader := bufio.NewReader(f)
	result := make([]byte, 5000)
	buffer := make([]byte, 4096)
	for {
		bureader.Read(buffer)
		copy(result, buffer)
		fmt.Println(string(result))
		if result[4095] != byte(32) {
			complementBuf, err := bureader.ReadBytes(byte(32))
			if err != nil {
				fmt.Println(err)
				return
			}
			copy(result[4096:], complementBuf)
		}
	}

	fmt.Println(string(result))
}
