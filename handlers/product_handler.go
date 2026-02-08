package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/heru-oktafian/tugas-02/models"
	"github.com/heru-oktafian/tugas-02/services"
	"github.com/heru-oktafian/tugas-02/tools"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProducts - GET /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		tools.JSONResponseNoData(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Products retrieved successfully", products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Failed to create product")
		return
	}

	tools.JSONResponse(w, http.StatusCreated, "Product created successfully", product)
}

// HandleProductByID - GET/PUT/DELETE /api/products/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		tools.JSONResponseNoData(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// GetByID - GET /api/products/{id}
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusNotFound, "Product not found")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Product retrieved successfully", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Failed to update product")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Product updated successfully", product)
}

// Delete - DELETE /api/products/{id}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponseNoData(w, http.StatusOK, "Product deleted successfully")
}
