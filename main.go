package main

import (
	"context"
	"fmt"

	pb "github.com/mariogao/microservice/proto"

	"github.com/micro/go-micro/v2"
)

func main() {
	//构建微服务
	server := micro.NewService()
	server.Init()
	// 调用服务端 /microServer/proto/hello.pb.micro.go中的方法
	// 得到一个结构体(该结构体实现了 client api 的方法)
	helloService := pb.NewHelloService("myMicro", server.Client())
	// 调用服务端 /microServer/proto/hello.pb.go中的方法
	// 得到入参结构体
	HelloReq := &pb.HelloReq{
		Name: "my name",
	}
	// 调用 /microServer/proto/hello.pb.micro.go中的方法(client api)
	// (client api 这个方法的具体实现中调用了 Server API 的MyHelloService 方法)
	// Server API 中的这个方法最终是由/microServer/controller中的 HelloService 实现了
	res, err := helloService.MyHelloService(context.Background(), HelloReq)
	fmt.Println("res:", res, "err:", err)

}
