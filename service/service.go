package service

import (
	"log"
	c "github.com/freundallein/color-gonverter/converter"
	pb "github.com/freundallein/color-gonverter/protobuf"
	"io"
)

type Gonverter struct{}

func (g *Gonverter) ConvertARGB_RGBA32(stream pb.GOnverter_ConvertARGB_RGBA32Server) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received ARGB: 0x%X\n", in.Argb)
		converted := c.Converter32(in.Argb)
		if err := stream.Send(&pb.RGBA32{Rgba: converted}); err != nil {
			return err
		}
		log.Printf("Sent RGBA: 0x%X\n", converted)

	}
}

func (g *Gonverter) ConvertARGB_RGBA64(stream pb.GOnverter_ConvertARGB_RGBA64Server) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received 0x%X\n", in.Argb)
		converted := c.Converter64(in.Argb)
		if err := stream.Send(&pb.RGBA64{Rgba: converted}); err != nil {
			return err
		}
		log.Printf("Sent 0x%X\n", converted)

	}
}
