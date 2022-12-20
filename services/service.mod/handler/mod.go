package handler

import (
	"context"
	"log"

	"github.com/mxbikes/mxbikesclient/services/service.mod/models"
	"github.com/mxbikes/mxbikesclient/services/service.mod/repository"
	protobuffer "github.com/mxbikes/protobuf/mod"
)

type Mod struct {
	protobuffer.UnimplementedModServiceServer
	postgres repository.ModPostgresRepository
}

// Return a new handler
func New(postgres repository.ModPostgresRepository) *Mod {
	return &Mod{postgres: postgres}
}

func (e *Mod) GetModByID(ctx context.Context, req *protobuffer.GetModByIDRequest) (*protobuffer.GetModByIDResponse, error) {
	log.Print("Received Mod.GetModByID request")

	mod, err := e.postgres.GetModByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &protobuffer.GetModByIDResponse{Mod: models.ModToProto(mod)}, nil
}

func (e *Mod) SearchMod(ctx context.Context, req *protobuffer.SearchModRequest) (*protobuffer.SearchModResponse, error) {
	log.Print("Received Mod.GetMods request")

	req.SearchText = "%" + req.SearchText + "%"

	if req.Size == 0 {
		req.Size = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := &models.ListQuery{
		Size: int(req.Size),
		Page: int(req.Page),
	}

	listQuery, err := e.postgres.SearchMod(ctx, req.SearchText, req.ModTypeCategoryIDs, query)
	if err != nil {
		return nil, err
	}

	var hasMore = false
	if listQuery.TotalPage > listQuery.Page {
		hasMore = true
	}

	var pagination = &protobuffer.Pagination{
		TotalCount: listQuery.TotalItems,
		TotalPages: int64(listQuery.TotalPage),
		Page:       int64(listQuery.Page),
		Size:       int64(listQuery.Size),
		HasMore:    hasMore,
	}

	return &protobuffer.SearchModResponse{Mods: models.ModsToProto(listQuery.Items), Pagination: pagination}, nil
}
