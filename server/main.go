package main

import (
  "context"
  "fmt"
	"time"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"

	proto "github.com/jperkelens/micro-path-example/proto"
)

type helloHandler struct {}

func newHelloHandler() proto.HelloHandler {
  return &helloHandler{}
}

func (handler *helloHandler) Greet(ctx context.Context, request *proto.GreetRequest, response *proto.GreetResponse) error {
  response.Message = "Hey how ya doin?"
  return nil
}

func main() {
  handle := newHelloHandler()
	grpcSvc := grpc.NewService(
		micro.Name("hello"),
		micro.RegisterTTL(time.Duration(30)*time.Second),
		micro.RegisterInterval(time.Duration(30)*time.Second),
	)

	grpcSvc.Init()
	proto.RegisterHelloHandler(grpcSvc.Server(), handle)

  err := grpcSvc.Run()
  if err != nil {
    fmt.Println(err.Error())
  }
}
