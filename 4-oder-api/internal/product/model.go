package product

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null;check:price >= 0"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}
