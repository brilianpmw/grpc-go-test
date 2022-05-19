package main

import (
	"context"
	"log"

	pb "github.com/brilianpmw/go-grpc-test/calculator/proto"
)

func (*Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.CalculatorResponse{Result:  in.Number1}, nil
}