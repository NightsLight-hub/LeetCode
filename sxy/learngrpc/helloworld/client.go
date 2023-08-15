/*
@Time : 2021/7/21 10:55
@Author : sunxy
@File : client
@description:
*/
package helloworld

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"nightslight.github.com/sxy/learngrpc/protobuf_introduction"
	"time"
)

func client() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
		return
	}
	client := protobuf_introduction.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatalln(err)
		return
	}
	go func() {
		for {
			if err := stream.Send(&protobuf_introduction.String{Value: "hi"}); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalln(err)
			}
			log.Println(reply.GetValue())
		}
	}()

	time.Sleep(100 * time.Second)

	//reply, err:=client.Hello(context.Background(), &protobuf_introduction.String{Value: "hello"})
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//log.Println(reply)
}
