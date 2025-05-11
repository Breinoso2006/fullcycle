package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/dto"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/entity"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/infra/database"
	entityPkg "github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create a new product
// @Description Create a new product with the given name and price
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductInput true "Product Request"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: "Id não pode ser vazio",
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Param product body dto.CreateProductInput true "Product Request"
// @Success 200
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Verifica se o ID foi passado na URL
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: "Id não pode ser vazio",
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	// Para garantir que o ID é válido, tentamos fazer o parse dele
	_, err := entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	// Verifica se o produto existe
	existingProduct, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	// Faz o parse do corpo da requisição para o struct Product
	var updatedProduct entity.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	// Mantém o CreatedAt e o ID do produto existente e atualiza os outros campos
	updatedProduct.CreatedAt = existingProduct.CreatedAt
	updatedProduct.ID = existingProduct.ID
	err = h.ProductDB.Update(&updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Failure 400 {object} Error
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: "Id não pode ser vazio",
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	err := h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ListProducts godoc
// @Summary List all products
// @Description List all products with pagination and sorting
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param limit query string false "Page size"
// @Param sort query string false "Sort by asc or desc"
// @Success 200 {array} entity.Product
// @Failure 404
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")
	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
