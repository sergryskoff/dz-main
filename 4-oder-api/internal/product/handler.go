// package product
package product

import (
	"fmt"
	"net/http"
	"order-api/pkg/middleware"
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
			middleware.LogError("Bad request error", map[string]interface{}{
				"error": err.Error(),
				"path":  r.URL.Path,
			})
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
			middleware.LogWarn("Product already exists", map[string]interface{}{
				"name": productRequst.Name,
				"path": r.URL.Path,
			})
			return
		}

		createProduct, err := l.Repo.Create(&product)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			middleware.LogError("Error creating product", map[string]interface{}{
				"error": err.Error(),
				"path":  r.URL.Path,
			})
			return
		}
		middleware.LogInfo("Product created successfully", map[string]interface{}{
			"product_id": createProduct.ID,
			"path":       r.URL.Path,
		})
		res.Json(w, createProduct, http.StatusCreated)
	}
}

func (l *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productRequst, err := req.HandleBody[ProductRequest](&w, r)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			middleware.LogError("Bad request error", map[string]interface{}{
				"error": err.Error(),
				"path":  r.URL.Path,
			})
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			middleware.LogError("Error parsing ID", map[string]interface{}{
				"error": err.Error(),
				"id":    idString,
				"path":  r.URL.Path,
			})
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
			middleware.LogError("Error updating product", map[string]interface{}{
				"error": err.Error(),
				"id":    id,
				"path":  r.URL.Path,
			})
			return
		}
		middleware.LogInfo("Product updated successfully", map[string]interface{}{
			"product_id": productUpdated.ID,
			"path":       r.URL.Path,
		})
		res.Json(w, productUpdated, http.StatusCreated)
	}
}

func (l *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			middleware.LogError("Error parsing ID", map[string]interface{}{
				"error": err.Error(),
				"id":    idString,
				"path":  r.URL.Path,
			})
			return
		}
		_, err = l.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, "Product not exist", http.StatusNotFound)
			middleware.LogWarn("Product not found", map[string]interface{}{
				"id":   id,
				"path": r.URL.Path,
			})
			return
		}
		err = l.Repo.Delete(uint(id))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			middleware.LogError("Error deleting product", map[string]interface{}{
				"error": err.Error(),
				"id":    id,
				"path":  r.URL.Path,
			})
			return
		}
		info := fmt.Sprintf("Delete product with ID: %d", id)
		middleware.LogInfo("Product deleted successfully", map[string]interface{}{
			"id":   id,
			"path": r.URL.Path,
		})
		res.Json(w, info, http.StatusOK)
	}
}

func (l *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			middleware.LogError("Error parsing ID", map[string]interface{}{
				"error": err.Error(),
				"id":    idString,
				"path":  r.URL.Path,
			})
			return
		}
		product, err := l.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, "Product not exist", http.StatusNotFound)
			middleware.LogWarn("Product not found", map[string]interface{}{
				"id":   id,
				"path": r.URL.Path,
			})
			return
		}
		middleware.LogInfo("Product retrieved successfully", map[string]interface{}{
			"id":   id,
			"path": r.URL.Path,
		})
		res.Json(w, product, http.StatusOK)
	}
}

func (l *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := l.Repo.GetAll()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			middleware.LogError("Failed to fetch products", map[string]interface{}{
				"error": err.Error(),
				"path":  r.URL.Path,
			})
			return
		}
		middleware.LogInfo("Products retrieved successfully", map[string]interface{}{
			"count": len(products),
			"path":  r.URL.Path,
		})
		res.Json(w, products, http.StatusOK)
	}
}
