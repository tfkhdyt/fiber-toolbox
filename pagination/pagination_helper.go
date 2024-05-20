package pagination

import "math"

// Meta struct represents metadata for paginated data.
type Meta struct {
	Page       uint `json:"page"`
	PageSize   uint `json:"page_size"`
	TotalItems uint `json:"total_items"`
	TotalPages uint `json:"total_pages"`
}

// NewMeta is a constructor function that creates and returns a new Meta struct.
// It takes the current page number, page size, and total number of items as arguments.
func NewMeta(page, pageSize, totalItems uint) *Meta {
	return &Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: uint(math.Ceil(float64(totalItems) / float64(pageSize))),
	}
}
