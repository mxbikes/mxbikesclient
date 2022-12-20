package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mxbikes/mxbikesclient/services/service.modType/handler"
	"github.com/mxbikes/mxbikesclient/services/service.modType/repository"
	protobuffer "github.com/mxbikes/protobuf/modType"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logLevel    = getEnv("LOG_LEVEL", "info")
	port        = getEnv("PORT", "localhost:4093")
	postgresUrl = getEnv("POSTGRES_URI", "host=localhost port=5432 user=postgres password=password sslmode=disable timezone=UTC connect_timeout=5")
)

func main() {
	/* Database */
	postgresConn, err := pgx.Connect(context.Background(), postgresUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to postgres database: %v\n", err)
		os.Exit(1)
	}
	defer postgresConn.Close(context.Background())
	postgresRepository := repository.NewPostgresRepository(postgresConn)

	/* Server */
	// Create a tcp listner
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to create listener", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	protobuffer.RegisterModTypeServiceServer(grpcServer, handler.New(postgresRepository))
	reflection.Register(grpcServer)

	// Start grpc server on listener
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
