package main

import (
	"log"
	"net"
	pb "github.com/brilianpmw/go-grpc-test/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}
func main() {
	var lis, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("falied to listen")
	}
	log.Printf("run  ")

	s:=grpc.NewServer()
	// fmt.Println(&Server{})
	pb.RegisterGreetServiceServer(s,&Server{})
	if err = s.Serve(lis) ; err != nil {
		log.Fatalf("failed load server")
	}
}
