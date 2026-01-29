package tools

import (
	"encoding/json"
	"net/http"

	"github.com/heru-oktafian/tugas-02/models"
)

// getStatusText returns "success" or "error" based on the HTTP status code
func getStatusText(status int) string {
	if status >= 200 && status < 300 {
		return "success"
	}
	return "error"
}

// JSONResponse sends a standard JSON response format / structure
func JSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := models.Response{
		Status:  getStatusText(status),
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

// JSONResponse sends a standard JSON response format / structure
func JSONResponseNoData(w http.ResponseWriter, status int, message string) {
	resp := models.ResponseNoData{
		Status:  getStatusText(status),
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
