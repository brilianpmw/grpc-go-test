package main

import (
	"log"
	"context"

	pb "github.com/brilianpmw/go-grpc-test/greet/proto"
	// "google.golang.org/grpc/credentials"


)


func doGreet(c pb.GreetServiceClient) {
	log.Printf("do greet jalan ")
	res,err := c.Greet(context.Background(),&pb.GreetRequest{
		Firstname : "Brili",
	})

	if(err != nil){
		log.Fatalf("do greet error",err)
	}

	log.Printf("result : %s ",res.Result)
}