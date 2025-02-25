package main

import (
	"gw_exchanger/internal/service"
	"gw_exchanger/internal/storages/postgres"
	"log"
	"net"

	exchange "github.com/apelsinkoo09/proto-exchange"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	storage := &postgres.StorageConn{DB: db}
	exchangeService := service.NewExchangeService(storage)

	server := grpc.NewServer()
	exchange.RegisterExchangeServiceServer(server, exchangeService)

<<<<<<< HEAD
	serv, err := net.Listen("tcp", "exchanger:50051")
=======
	serv, err := net.Listen("tcp", ":50051")
>>>>>>> e5f495f26039e46d21c26b8e0d9fd5f9696ba28f
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("gRPC server is running on port 50051")

	if err := server.Serve(serv); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
