package helpers

import (
	"AutomobilesAPI/dtos/response"
	"encoding/json"
	"net/http"
)

func IsMethodsAllowed(method, httpmethod string, w http.ResponseWriter) bool {
	if method != httpmethod {
		reposen := response.ApiResponse{
			Code:    http.StatusMethodNotAllowed,
			Status:  "NotAllowed",
			Message: "Method not allowed",
		}

		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(reposen)
		IsPanicErrors(err)
		return false
	}
	return true
}
