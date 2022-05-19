
package main

import (
	"log"
	pb "github.com/brilianpmw/go-grpc-test/calculator/proto"
	// "google.golang.org/grpc/credentials"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

var addr string = "localhost:50052"


func main() {
	var conn,err = grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error when create dial ",err)
	}
	defer conn.Close()

	 c:=pb.NewCalculatorServiceClient(conn)

	 doCalculate(c)

}

