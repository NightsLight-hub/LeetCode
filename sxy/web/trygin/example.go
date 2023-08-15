/*
@Time : 2021/6/21 9:22
@Author : sunxy
@File : example
@description:
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		s, _ := ioutil.ReadAll(c.Request.Body) //把	body 内容读入字符串 s
		fmt.Printf("%s\n", s)
		c.JSON(200, string(s))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
