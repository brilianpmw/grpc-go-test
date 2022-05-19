package main

import (
	"log"
	"context"

	pb "github.com/brilianpmw/go-grpc-test/calculator/proto"
	// "google.golang.org/grpc/credentials"


)


func doCalculate(c pb.CalculatorServiceClient) {
	log.Printf("do calculator jalan ")
	res,err := c.Greet(context.Background(),&pb.CalculatorRequest{
		Number1 : 1,
		Number2:2,
	})

	if(err != nil){
		log.Fatalf("do greet error",err)
	}

	log.Printf("result : %s ",res)
}