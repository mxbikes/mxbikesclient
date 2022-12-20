package handler

import (
	"context"
	"log"

	"github.com/mxbikes/mxbikesclient/services/service.modTypeCategory/models"
	"github.com/mxbikes/mxbikesclient/services/service.modTypeCategory/repository"
	protobuffer "github.com/mxbikes/protobuf/modTypeCategory"
)

type ModTypeCategory struct {
	protobuffer.UnimplementedModTypeCategoryServiceServer
	postgres repository.ModTypeCategoryPostgresRepository
}

// Return a new handler
func New(postgres repository.ModTypeCategoryPostgresRepository) *ModTypeCategory {
	return &ModTypeCategory{postgres: postgres}
}

func (e *ModTypeCategory) GetModTypeCategoryByID(ctx context.Context, req *protobuffer.GetModTypeCategoryByIDRequest) (*protobuffer.GetModTypeCategoryByIDResponse, error) {
	log.Print("Received ModTypeCategory.GetModTypeCategoryByID request")

	modTypeCategory, err := e.postgres.GetModTypeCategoryByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModTypeCategoryByIDResponse{ModTypeCategory: models.ModTypeCategoryToProto(modTypeCategory)}, nil
}

func (e *ModTypeCategory) GetModTypeCategoriesByModTypeID(ctx context.Context, req *protobuffer.GetModTypeCategoriesByModTypeIDRequest) (*protobuffer.GetModTypeCategoriesByModTypeIDResponse, error) {
	log.Print("Received ModTypeCategory.GetModTypeCategories request")

	modTypeCategories, err := e.postgres.GetModTypeCategoriesByModTypeID(ctx, req.ModTypeID)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModTypeCategoriesByModTypeIDResponse{ModTypeCategories: models.ModTypeCategoriesToProto(modTypeCategories)}, nil
}

func (e *ModTypeCategory) GetModTypeCategories(ctx context.Context, req *protobuffer.GetModTypeCategoriesRequest) (*protobuffer.GetModTypeCategoriesResponse, error) {
	log.Print("Received ModTypeCategory.GetModTypeCategories request")

	modTypeCategories, err := e.postgres.GetAllModTypeCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModTypeCategoriesResponse{ModTypeCategories: models.ModTypeCategoriesToProto(modTypeCategories)}, nil
}
