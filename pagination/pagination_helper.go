package pagination

import "math"

type Meta struct {
	Page       uint `json:"page"`
	PageSize   uint `json:"page_size"`
	TotalItems uint `json:"total_items"`
	TotalPages uint `json:"total_pages"`
}

func NewMeta(page, pageSize, totalItems uint) *Meta {
	return &Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: uint(math.Ceil(float64(totalItems) / float64(pageSize))),
	}
}
