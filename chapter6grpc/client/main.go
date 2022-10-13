package main

import (
	"context"
	"fmt"
	pb2 "gocode/mygo2/chapter6grpc/pb"
	"log"

	"google.golang.org/grpc"
)
const (
	// Address 连接地址
	Address string = ":8000"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	//// 建立gRPC连接
	//grpcClient := pb.NewSimpleClient(conn)
	//// 创建发送结构体
	//s := "grpc"
	//req := pb.SimpleRequest{
	//	Data: s,
	//}
	//// 调用我们的服务(Route方法)
	//// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	//res, err := grpcClient.Route(context.Background(), &req)
	//if err != nil {
	//	log.Fatalf("Call Route err: %v", err)
	//}
	grpcClient := pb2.NewSimpleClient(conn)

	req := pb2.SumRequest{
		Number1: 1,
		Number2: 4,
	}
	res, err := grpcClient.GetSum(context.Background(), &req)
	if err != nil {
		return
	}
	//res.Reset()
	// 打印返回值
	fmt.Println(res.GetSum())

	log.Println(res)
}