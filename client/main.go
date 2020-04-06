package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	firstGRPC "grpc-demo/pb"
	"os"
	"time"
)

const (
	serverAddress = ":8091"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	gRPCClient := firstGRPC.NewFirstGRPCClient(conn)

	// Contact the server and print out its response.
	paramName := "defaultName"
	paramValue := "defaultValue"
	if len(os.Args) > 2 {
		paramName = os.Args[1]
		paramValue = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	gRPCResponse, err := gRPCClient.TestGRPC(ctx, &firstGRPC.RequestParam{Name: paramName, Value: paramValue})
	if err != nil {
		fmt.Println("call TestGRPC:", err)
		return
	}
	fmt.Println("call TestGRPC res:", gRPCResponse.GetMessage())
}
