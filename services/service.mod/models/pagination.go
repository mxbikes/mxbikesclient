package models

import (
	"math"
)

const (
	defaultSize = 10
	defaultPage = 1
)

type ListResult[T any] struct {
	Size       int   `json:"size,omitempty" bson:"size"`
	Page       int   `json:"page,omitempty" bson:"page"`
	TotalItems int64 `json:"totalItems,omitempty" bson:"totalItems"`
	TotalPage  int   `json:"totalPage,omitempty" bson:"totalPage"`
	Items      []T   `json:"items,omitempty" bson:"items"`
}

func NewListResult[T any](items []T, size int, page int, totalItems int64) *ListResult[T] {
	listResult := &ListResult[T]{Items: items, Size: size, Page: page, TotalItems: totalItems}

	listResult.TotalPage = getTotalPages(totalItems, size)

	return listResult
}

// GetTotalPages Get total pages int
func getTotalPages(totalCount int64, size int) int {
	d := float64(totalCount) / float64(size)
	return int(math.Ceil(d))
}

type ListQuery struct {
	Size    int    `query:"size" json:"size,omitempty"`
	Page    int    `query:"page" json:"page,omitempty"`
	OrderBy string `query:"orderBy" json:"orderBy,omitempty"`
}

// GetOffset Get offset
func (q *ListQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// GetLimit Get limit
func (q *ListQuery) GetLimit() int {
	return q.Size
}

// GetOrderBy Get OrderBy
func (q *ListQuery) GetOrderBy() string {
	return q.OrderBy
}

// GetPage Get OrderBy
func (q *ListQuery) GetPage() int {
	return q.Page
}

// GetSize Get OrderBy
func (q *ListQuery) GetSize() int {
	return q.Size
}
