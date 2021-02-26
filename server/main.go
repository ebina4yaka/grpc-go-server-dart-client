package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ebina4yaka/grpc-go-dart/pokemon"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {

	fmt.Println("Go gRPC Pokemon Server!")

	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Listen to :%s\n", port)

	pokemonServer := pokemon.Server{}

	grpcServer := grpc.NewServer()

	pokemon.RegisterPokemonServiceServer(grpcServer, &pokemonServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
