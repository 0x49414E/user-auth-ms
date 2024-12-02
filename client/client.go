package client

import (
	"context"
	"fmt"
	"log"
	"user_auth/pb"

	"google.golang.org/grpc"
)

const usuario = "JorgeElPanzon"
const password = "123456"

func RunClient() {
	// Set up a connection to the server.

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	// Test the Register method
	registerResponse, err := client.Register(context.Background(), &pb.RegisterRequest{
		Username: usuario,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	fmt.Printf("Register Response: %s\n", registerResponse.Message)

	// Test the Login method
	loginResponse, err := client.Login(context.Background(), &pb.LoginRequest{
		Username: usuario,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	fmt.Printf("Login Response: %s\n", loginResponse.Token)

	//// Test the UpdateUserDetails method
	//updateUserDetailsResponse, err := client.UpdateUserDetails(context.Background(), &pb.UpdateUserDetailsRequest{
	//	Id:         1,
	//	Username:   "testuser",
	//	Name:       "Test",
	//	Lastname:   "User",
	//	Dni:        12345678,
	//	Address:    "Test Address",
	//	PostalCode: 12345,
	//	Phone:      123456789,
	//})
	//if err != nil {
	//	log.Fatalf("could not update user details: %v", err)
	//}
	//fmt.Printf("Update User Details Response: %s\n", updateUserDetailsResponse.Message)

}
