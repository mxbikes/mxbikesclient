package handler

import (
	"context"
	"log"

	"github.com/mxbikes/mxbikesclient/services/service.modType/models"
	"github.com/mxbikes/mxbikesclient/services/service.modType/repository"
	protobuffer "github.com/mxbikes/protobuf/modType"
)

type ModType struct {
	protobuffer.UnimplementedModTypeServiceServer
	postgres repository.ModTypePostgresRepository
}

// Return a new handler
func New(postgres repository.ModTypePostgresRepository) *ModType {
	return &ModType{postgres: postgres}
}

func (e *ModType) GetModTypeByID(ctx context.Context, req *protobuffer.GetModTypeByIDRequest) (*protobuffer.GetModTypeByIDResponse, error) {
	log.Print("Received ModType.GetModTypeByID request")

	modType, err := e.postgres.GetModTypeByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModTypeByIDResponse{ModType: models.ModTypeToProto(modType)}, nil
}

func (e *ModType) GetModTypes(ctx context.Context, req *protobuffer.GetModTypesRequest) (*protobuffer.GetModTypesResponse, error) {
	log.Print("Received ModType.GetModTypes request")

	modTypes, err := e.postgres.GetAllModTypes(ctx)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModTypesResponse{ModTypes: models.ModTypesToProto(modTypes)}, nil
}
