/*
@Time : 2021/7/21 10:50
@Author : sunxy
@File : helloworld
@description:
*/
package helloworld

import (
	"context"
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
	"nightslight.github.com/sxy/learngrpc/protobuf_introduction"
)

const ServiceName = "HelloWorld"

type HelloWorld struct {
}

func (h *HelloWorld) Hello(requset protobuf_introduction.String, reply *protobuf_introduction.String) error {
	reply.Value = "heelo" + requset.GetValue()
	return nil
}

type IHelloWorld interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc IHelloWorld) error {
	return rpc.RegisterName(ServiceName, svc)
}

type HelloWorldRpc struct {
	*rpc.Client
}

func (h *HelloWorldRpc) DialRpc(network, address string) error {
	c, err := jsonrpc.Dial(network, address)
	if err != nil {
		return err
	}
	h.Client = c
	return nil
}

func (h *HelloWorldRpc) hello(request string, reply *string) error {
	return h.Client.Call(ServiceName+".Hello", request, reply)
}

type HelloServiceImpl struct {
}


func (h HelloServiceImpl) Channel(stream protobuf_introduction.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		reply := &protobuf_introduction.String{Value: "hello" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (h HelloServiceImpl) Hello(ctx context.Context, s *protobuf_introduction.String) (*protobuf_introduction.String, error) {
	return &protobuf_introduction.String{
		Value: "hello" + s.GetValue(),
	}, nil
}
