package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	protobuffer_mod "github.com/mxbikes/protobuf/mod"
	protobuffer_modImage "github.com/mxbikes/protobuf/modImage"
	protobuffer_modType "github.com/mxbikes/protobuf/modType"
	protobuffer_modTypeCategory "github.com/mxbikes/protobuf/modTypeCategory"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	port                      = getEnv("PORT", "localhost:4105")
	URLServiceMod             = getEnv("SERVICE_MOD_URL", "localhost:4089")
	URLServiceModImage        = getEnv("SERVICE_MODIMAGE_URL", "localhost:4091")
	URLServiceModType         = getEnv("SERVICE_MODTYPE_URL", "localhost:4093")
	URLServiceModTypeCategory = getEnv("SERVICE_MODTYPECATEGORY_URL", "localhost:4096")
)

func main() {
	gwmux := runtime.NewServeMux()

	// Register ServiceMod
	err := protobuffer_mod.RegisterModServiceHandlerFromEndpoint(context.Background(), gwmux, URLServiceMod, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalln("Failed to register ServiceMod:", err)
	}

	// Register ServiceModImage
	err = protobuffer_modImage.RegisterModImageServiceHandlerFromEndpoint(context.Background(), gwmux, URLServiceModImage, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalln("Failed to register ServiceModImage:", err)
	}

	// Register ModType
	err = protobuffer_modType.RegisterModTypeServiceHandlerFromEndpoint(context.Background(), gwmux, URLServiceModType, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalln("Failed to register ModType:", err)
	}

	// Register ModTypeCategory
	err = protobuffer_modTypeCategory.RegisterModTypeCategoryServiceHandlerFromEndpoint(context.Background(), gwmux, URLServiceModTypeCategory, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalln("Failed to register ModTypeCategory:", err)
	}

	gwServer := &http.Server{
		Addr:    port,
		Handler: cors(gwmux),
	}

	log.Print("Serving gRPC-Gateway on http://" + port)
	log.Fatalln(gwServer.ListenAndServe())
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)

	})
}
