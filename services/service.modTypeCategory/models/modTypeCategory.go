package models

import (
	protobuffer "github.com/mxbikes/protobuf/modTypeCategory"
)

type ModTypeCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ModTypeID string `json:"id_modType"`
}

func ModTypeCategoryToProto(modTypeCategory *ModTypeCategory) *protobuffer.ModTypeCategory {
	return &protobuffer.ModTypeCategory{
		ID:        modTypeCategory.ID,
		Name:      modTypeCategory.Name,
		ModTypeID: modTypeCategory.ModTypeID,
	}
}

func ModTypeCategoriesToProto(modTypeCategories []*ModTypeCategory) []*protobuffer.ModTypeCategory {
	orders := make([]*protobuffer.ModTypeCategory, 0, len(modTypeCategories))
	for _, projection := range modTypeCategories {
		orders = append(orders, ModTypeCategoryToProto(projection))
	}
	return orders
}
