package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	firstGRPC "grpc-demo/pb"
	"net"
)

type FirstGRPCServer struct {
	firstGRPC.UnimplementedFirstGRPCServer
}

func (firstGRPCServer *FirstGRPCServer) TestGRPC(ctx context.Context, requestParam *firstGRPC.RequestParam) (
	*firstGRPC.ReplyBody, error) {
	fmt.Println("requestParam:", requestParam.String())
	return &firstGRPC.ReplyBody{Code: 0, Message: "first gRPC call success"}, nil
}

func main() {
	listenAddress, err := net.Listen("tcp", "127.0.0.1:8091")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	fmt.Println("listen success")

	gRPCServer := grpc.NewServer()
	firstGRPC.RegisterFirstGRPCServer(gRPCServer, &FirstGRPCServer{})
	if err := gRPCServer.Serve(listenAddress); err != nil {
		fmt.Println("failed to server:", err)
		return
	}
}
