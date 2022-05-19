package main

import (
	"log"
	pb "github.com/brilianpmw/go-grpc-test/greet/proto"
	// "google.golang.org/grpc/credentials"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

var addr string = "localhost:50051"


func main() {
	var conn,err = grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error when create dial ",err)
	}
	defer conn.Close()

	 c:=pb.NewGreetServiceClient(conn)

	doGreet(c)

}