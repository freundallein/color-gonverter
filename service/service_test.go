package service

import (
	pb "github.com/freundallein/color-gonverter/protobuf"
	"google.golang.org/grpc"
	"testing"
)

func TestService(t *testing.T) {
	// Check if service implementation mathces proto definition.
	server := grpc.NewServer()
	pb.RegisterGOnverterServer(server, &Gonverter{})
	server.Stop()
}
