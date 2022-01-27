package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"pcbook-go/pb"
	"pcbook-go/sample"
	"time"
)

func main() {
	fmt.Println("Hello, world! from Client")
	serviceAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dialsing server %s\n", *serviceAddress)

	conn, err := grpc.Dial(*serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't dial server'", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// send request to server
	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Printf("create laptop with id: %s", res.Id)
}
