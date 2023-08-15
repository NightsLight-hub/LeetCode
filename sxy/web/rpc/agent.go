/*
@Time : 2021/7/8 16:53
@Author : sunxy
@File : agent
@description:
*/
package rpc

import (
	"fmt"
	"log"
	"net/rpc"
)

func agent() {
	client, err := rpc.Dial("tcp", "localhost:12234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
