/*
@Time : 2021/7/21 10:49
@Author : sunxy
@File : server.go
@description:
*/
package helloworld

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"nightslight.github.com/sxy/learngrpc/protobuf_introduction"
)

func server() {
	grpcServer := grpc.NewServer()
	protobuf_introduction.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	list, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
