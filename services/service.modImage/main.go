package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mxbikes/mxbikesclient/services/service.modImage/handler"
	"github.com/mxbikes/mxbikesclient/services/service.modImage/repository"
	protobuffer "github.com/mxbikes/protobuf/modImage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logLevel       = getEnv("LOG_LEVEL", "info")
	port           = getEnv("PORT", "localhost:4091")
	postgresUrl    = getEnv("POSTGRES_URI", "host=localhost port=5432 user=postgres password=password sslmodImagee=disable timezone=UTC connect_timeout=5")
	minioHost      = getEnv("MINIO_HOST", "localhost:9001")
	minioAccessKey = getEnv("MINIO_ACCESKEY", "NuVwLvC17qn3RuXs")
	minioSecret    = getEnv("MINIO_SECRET", "eVdFBNNJd1pr40K5TkkaebsKqs4q97VF")
)

func main() {
	/* Database */
	minioConn, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("LgkBiLyczoKKdq4c", "ECGn9urAXueoJnm58WRg7dg7iSEkCS3n", ""),
		Secure: false,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to minio database: %v\n", err)
		os.Exit(1)
	}
	minioRepository := repository.NewMinioRepository(minioConn)

	/* Server */
	// Create a tcp listner
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to create listener", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	protobuffer.RegisterModImageServiceServer(grpcServer, handler.New(minioRepository))
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
