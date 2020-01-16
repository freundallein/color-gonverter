package main

import (
	"context"
	"fmt"
	pb "github.com/freundallein/color-gonverter/protobuf"
	"google.golang.org/grpc"
	"io"
	"os"
)

// test data (2 points in 1 uint64)
var data = []*pb.ARGB64{
	&pb.ARGB64{Argb: 0xAABBCCDDAABBCCDD},
	&pb.ARGB64{Argb: 0x1122334411223344},
	&pb.ARGB64{Argb: 0x5566778855667788},
	&pb.ARGB64{Argb: 0x9910EEFF9910EEFF},
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if conn != nil {
		defer conn.Close()
		fmt.Println("Connected to tcp:7777")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "did not connect: %v\n", err)
	}
	c := pb.NewGOnverterClient(conn)
	stream, err := c.ConvertARGB_RGBA64(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling ConvertARGB_RGBA32: %v\n", err)
	}
	for _, element := range data {
		if err := stream.Send(element); err != nil {
			fmt.Fprintf(os.Stderr, "failed to send: %v\n", err)
		}
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to receive: %v\n", err)
		}
		fmt.Printf("Converted 0x%X to 0x%X\n", element.Argb, in.Rgba)
	}
}
