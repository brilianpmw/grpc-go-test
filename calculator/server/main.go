package main

import (
	"log"
	"net"
	pb "github.com/brilianpmw/go-grpc-test/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}
func main() {
	var lis, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("falied to listen")
	}
	log.Printf("run calculator ")

	s:=grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s,&Server{})
	if err = s.Serve(lis) ; err != nil {
		log.Fatalf("failed load server")
	}
}
