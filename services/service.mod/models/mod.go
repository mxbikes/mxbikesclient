package models

import (
	"time"

	"github.com/google/uuid"
	protobuffer "github.com/mxbikes/protobuf/mod"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Mod struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	ModTypeCategoryID uuid.UUID `json:"modTypeCategoryID"`

	ReleaseYear int16     `json:"releaseYear"`
	CreateAt    time.Time `json:"createAt"`
}

func ModToProto(mod *Mod) *protobuffer.Mod {
	return &protobuffer.Mod{
		ID:                mod.ID,
		Name:              mod.Name,
		Description:       mod.Description,
		ModTypeCategoryID: mod.ModTypeCategoryID.String(),
		ReleaseYear:       int32(mod.ReleaseYear),
		CreateAt:          timestamppb.New(mod.CreateAt),
	}
}

func ModsToProto(mods []*Mod) []*protobuffer.Mod {
	orders := make([]*protobuffer.Mod, 0, len(mods))
	for _, projection := range mods {
		orders = append(orders, ModToProto(projection))
	}
	return orders
}
