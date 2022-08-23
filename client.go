package main

import (
	"context"
	"fmt"
	pb "gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 1.链接服务端
// 2.实例化gRPC客户端
// 3.调用

func main() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接异常%s\n", err)
	}
	defer conn.Close()
	// 2.实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3.组装请求参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4.调用
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("相应异常%s\n", err)
	}
	fmt.Printf("相应结果%v\n", response)
}
