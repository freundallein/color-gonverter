package main

import (
	"fmt"
	"log"
	"net"
	"time"
	pb "github.com/freundallein/color-gonverter/protobuf"
	service "github.com/freundallein/color-gonverter/service"
	"google.golang.org/grpc"
)

const (
	timeFormat = "02.01.2006 15:04:05"
)

type logWriter struct {
}

// Write - custom logger formatting
func (writer logWriter) Write(bytes []byte) (int, error) {
	msg := fmt.Sprintf("%s | [converter] %s", time.Now().UTC().Format(timeFormat), string(bytes))
	return fmt.Print(msg)
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
	}
	log.Printf("Listen on tcp:%d\n", 7777)
	grpcServer := grpc.NewServer()
	pb.RegisterGOnverterServer(grpcServer, &service.Gonverter{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("failed to serve: %s", err)
	}
}
