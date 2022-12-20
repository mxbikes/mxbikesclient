package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mxbikes/mxbikesclient/services/service.modImage/models"
	"github.com/mxbikes/mxbikesclient/services/service.modImage/repository"
	protobuffer "github.com/mxbikes/protobuf/modImage"
)

type ModImage struct {
	protobuffer.UnimplementedModImageServiceServer
	minio repository.ModImageMinioRepository
}

// Return a new handler
func New(minio repository.ModImageMinioRepository) *ModImage {
	return &ModImage{minio: minio}
}

func (e *ModImage) GetModImagesByModID(ctx context.Context, req *protobuffer.GetModImagesByModIDRequest) (*protobuffer.GetModImagesByModIDResponse, error) {
	log.Print("Received ModImage.GetModImagesByModID request")

	modID, err := uuid.Parse(req.ModID)
	if err != nil {
		return nil, err
	}

	modImages, err := e.minio.GetModImagesByModID(ctx, modID.String())
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModImagesByModIDResponse{ModImage: models.ModImagesToProto(modImages)}, nil
}
