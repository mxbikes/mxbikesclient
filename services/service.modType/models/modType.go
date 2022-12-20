package models

import (
	protobuffer "github.com/mxbikes/protobuf/modType"
)

type ModType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ModTypeToProto(modType *ModType) *protobuffer.ModType {
	return &protobuffer.ModType{
		ID:   modType.ID,
		Name: modType.Name,
	}
}

func ModTypesToProto(modTypes []*ModType) []*protobuffer.ModType {
	orders := make([]*protobuffer.ModType, 0, len(modTypes))
	for _, projection := range modTypes {
		orders = append(orders, ModTypeToProto(projection))
	}
	return orders
}
