// package product
package product

import (
	"fmt"
	"log"
	"net/http"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	Repo *ProductRepository
}

func NewProductHandler(r *http.ServeMux, repo *ProductRepository) {
	h := &ProductHandler{
		Repo: repo,
	}

	r.HandleFunc("POST /products", h.Create())
	r.HandleFunc("PATCH /products/{id}", h.Update())
	r.HandleFunc("DELETE /products/{id}", h.Delete())
	r.HandleFunc("GET /products/{id}", h.GetProductById())
	r.HandleFunc("GET /products", h.GetAll())
}

func (l *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productRequst, err := req.HandleBody[ProductRequest](&w, r)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			log.Printf("Bad request error: %v\n", err)
			return
		}

		var product Product

		product = Product{
			Name:        productRequst.Name,
			Description: productRequst.Description,
			Price:       productRequst.Price,
			Images:      productRequst.Images,
		}

		existedProduct, _ := l.Repo.GetByName(productRequst.Name)
		if existedProduct != nil {
			res.Json(w, "Product exist", http.StatusFound)
			log.Printf("Product exist, Name %s\n", productRequst.Name)
			return
		}

		createProduct, err := l.Repo.Create(&product)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error creating product: %v\n", err)
			return
		}
		res.Json(w, createProduct, http.StatusCreated)
	}
}

func (l *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productRequst, err := req.HandleBody[ProductRequest](&w, r)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			log.Printf("Bad request error: %v\n", err)

			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error parse ID: %v\n", err)
			return
		}

		var product Product

		product = Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        productRequst.Name,
			Description: productRequst.Description,
			Price:       productRequst.Price,
			Images:      productRequst.Images,
		}

		productUpdated, err := l.Repo.Update(&product)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error updating product: %v\n", err)
			return
		}
		res.Json(w, productUpdated, http.StatusCreated)
	}
}

func (l *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error parse ID: %v\n", err)

			return
		}
		_, err = l.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, "Product not exist", http.StatusNotFound)
			log.Printf("Product not exist, ID %d\n", id)
			return
		}
		err = l.Repo.Delete(uint(id))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error deleting product: %v\n", err)
			return
		}
		info := fmt.Sprintf("Delete product with ID: %d", id)
		log.Printf("Delete product with ID: %d\n", id)
		res.Json(w, info, http.StatusOK)
	}
}

func (l *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error parse ID: %v\n", err)
			return
		}
		product, err := l.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, "Product not exist", http.StatusNotFound)
			log.Printf("Product not exist, ID %d\n", id)
			return
		}
		res.Json(w, product, http.StatusOK)
	}
}

func (l *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := l.Repo.GetAll()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			log.Printf("Failed to fetch products: %v\n", err)
			return
		}
		res.Json(w, products, http.StatusOK)
	}
}
