package response

import "net/http"

type ApiResponse struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// The function returns an ApiResponse with a Bad Request status and a message.
func GetBadRequest(message string) ApiResponse {
	return ApiResponse{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: message,
	}
}

// The function returns an ApiResponse with a status code of 200 and a message.
func OkRequest(message string) ApiResponse {
	return ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: message,
	}
}

// The function returns an ApiResponse with a 404 status code and a message for a not found request.
func NotFoundRequest(message string) ApiResponse {
	return ApiResponse{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: message,
	}
}
