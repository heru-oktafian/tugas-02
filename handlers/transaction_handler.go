package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/heru-oktafian/tugas-02/models"
	"github.com/heru-oktafian/tugas-02/services"
	"github.com/heru-oktafian/tugas-02/tools"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandlerCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		tools.JSONResponseNoData(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, "Failed to checkout")
		return
	}

	tools.JSONResponse(w, http.StatusOK, "Checkout successfully", transaction)
}

func (h *TransactionHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	report, err := h.service.GetReport(startDate, endDate)
	if err != nil {
		tools.JSONResponseNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}
