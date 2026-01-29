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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		tools.JSONResponseNoData(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Failed to create category")
		return
	}

	tools.JSONResponse(w, http.StatusCreated, "Category created successfully", category)
}

// HandleCategoryByID - GET/PUT/DELETE /api/categories/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

// GetByID - GET /api/categories/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to retrieve category")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Category retrieved successfully", category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Category updated successfully", category)
}

// Delete - DELETE /api/categories/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	tools.JSONResponseNoData(w, http.StatusOK, "Category deleted successfully")
}
