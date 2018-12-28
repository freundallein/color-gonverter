package main

import (
	pb "./protobuf"
	service "./service"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
	}
	fmt.Printf("Listen on tcp:%d\n", 7777)
	grpcServer := grpc.NewServer()
	pb.RegisterGOnverterServer(grpcServer, &service.Gonverter{})
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve: %s", err)
	}
}
