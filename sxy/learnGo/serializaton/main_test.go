/*
Package serializaton
@Time : 2023/10/7 10:11
@Author : sunxy
@File : main
@description:
*/
package serializaton

import (
	"encoding/json"
	"fmt"
	"testing"
)

type dd struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

// Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON object.

func Test_main(t *testing.T) {
	str := "EvcDChduXzE4MTY1Mjg3MDdfMjI4MTQwMjA5MxLbAwrYAwoKMTgxNjUyODcwNxIKMjI4MTQwMjA5MyIgNDVhZTY5NzRmNGEzYjI1YjYyZmE3ZGExMjc0NjVkMDcqIDY0ZDNmNjdiNTMzZGVhY2Y4OTg2ZGJhMTMwM2MzNTcxSAFQyAFYmApiZnsiZGV0YWlsIjoie1wiZnJvbVVzZXJJRFwiOlwiMTgxNjUyODcwN1wiLFwidG9Vc2VySURcIjpcIjIyODE0MDIwOTNcIixcInN0YXR1c1wiOjEsXCJwbGF0Zm9ybUlEXCI6M30ifXjhsM+/sDGAAeGwz7+wMZoBCwoHaGlzdG9yeRAAmgEWChJjb252ZXJzYXRpb25V\ncGRhdGUQAJoBDwoLb2ZmbGluZVB1c2gQAJoBDgoKcGVyc2lzdGVudBAAmgEcChhzZW5kZXJDb252ZXJzYXRpb25VcGRhdGUQAJoBGgoWc2VuZGVyTm90aWZpY2F0aW9uUHVzaBAAmgEPCgt1bnJlYWRDb3VudBAAmgEVChFpc05vdE5vdGlmaWNhdGlvbhAAmgEVChFyZWFjdGlvbkZyb21DYWNoZRAAmgEOCgpub3RQcml2YXRlEACaAQ0KCWlzU2VuZE1zZxAAmgEOCgpzZW5kZXJTeW5jEACiAQA="
	di1 := dd{
		Name: "di1",
		Data: []byte(str),
	}

	di1Bytes, err := json.Marshal(di1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", di1Bytes)
	// di2 := new(dd)
	// json.Unmarshal(di1Bytes, di2)
	// fmt.Printf("%#v\n", di2)
	// fmt.Printf("%s\n", di2.Data)
}
