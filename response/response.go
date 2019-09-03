package response

import (
	"encoding/json"
	"net/http"

	"github.com/adhistria/ijahstore/model"
)

func APISuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	var successResponse model.SuccessResponse
	successResponse.Data = data
	successResponse.Message = "Success"
	json.NewEncoder(w).Encode(successResponse)
}

func APIErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	var errorResponse model.ErrorResponse
	errorResponse.Message = message
	json.NewEncoder(w).Encode(errorResponse)
}
