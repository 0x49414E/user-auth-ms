package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"user_auth/client"
	"user_auth/handlers"
	"user_auth/internals"
	"user_auth/pb"
	"user_auth/repositories"
	"user_auth/services"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Config struct {
	JWTSecret string
	Addr      string
	Transport string
	DBUrl     string
}

func LoadConfig() Config {
	return Config{
		JWTSecret: os.Getenv("SECRET_KEY"), 
		Addr:      ":50051",
		Transport: "tcp",
		DBUrl:     os.Getenv("DB_URL"), 
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := LoadConfig()

	db, err := sql.Open("mysql", cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	jwtManager := internals.NewJWTManager(cfg.JWTSecret, time.Hour)
	authService := services.NewAuthService(userRepo, jwtManager)
	authHandler := handlers.NewAuthHandler(authService)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authHandler)

	listener, err := net.Listen(cfg.Transport, cfg.Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("gRPC server is running on port 50051")

	go client.RunClient()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

//test webhook
