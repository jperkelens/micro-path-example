package main

import (
	"fmt"
	"os"
	"time"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"

	proto "github.com/jperkelens/micro-example/proto"
)

type helloHandler struct {}

func newHelloHandler() proto.HelloHandler {
  return &helloHandler{}
}

func (handler *helloHandler) Greet(ctx contect.Context, request *proto.GreetRequest, response *proto.GreetResponse) error {
  response.Message = "Hey how ya doin?"
}

func main() {
	handle := handler.newHelloHandler()

	grpcSvc := grpc.NewService(
		micro.Name("hello"),
		micro.RegisterTTL(time.Duration(cfg.RegistryTTL)*time.Second),
		micro.RegisterInterval(time.Duration(cfg.RegistryInterval)*time.Second),
	)

	grpcSvc.Init()
	proto.RegisterHelloHandler(rpcSvc.Server(), handle)
  grpcSvc.Run()
}
