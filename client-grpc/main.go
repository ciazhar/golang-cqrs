package main

import (
	"context"
	"fmt"
	"github.com/ciazhar/golang-cqrs/api-grpc/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := api.NewSocialServiceClient(cc)
	request := &api.SocialRequest{Id: "Jeremy"}

	resp, _ := client.GetByID(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)
}
