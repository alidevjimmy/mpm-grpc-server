package main

import (
	"context"
	"esme_team/src/pbs/userpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("185.235.43.194:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()
	// Create CLient
	c := userpb.NewUserServiceClient(cc)
	request := &userpb.LoginRequest{
		Phone:    "09384720059",
		Password: "password",
	}
	gResponse, err := c.Login(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling Login: %v", err)
	}
	log.Printf("gRPC Response: %v", gResponse.User)
}
