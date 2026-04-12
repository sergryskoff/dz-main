package product

import "github.com/lib/pq"

type ProductRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Price       float64        `json:"price" validate:"required"`
	Images      pq.StringArray `json:"images"`
}
