package main

import (
	"context"
	_ "gw_exchanger/internal/storages/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ExchangeServer struct {
	pb.UnimplementedExchangeServiceServer
	storage Strorage
}

// Реализация методов gRPC
func (s *ExchangeServer) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.storage.GetAllExchangeRates()
	if err != nil {
		return nil, err
	}
	return &pb.ExchangeRatesResponse{Rates: rates}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterExchangeServiceServer(server, &ExchangeServer{})
	log.Println("Server is running on port :50051")
	server.Serve(lis)
}
