package helpers

import (
	"encoding/json"
	"net/http"
)

// The function reads the request body of an HTTP request and decodes it into a specified result
// interface, while also checking for any errors.
func ReadRequestBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body).Decode(&result)
	IsPanicErrors(decoder)
}

// The function writes a JSON response to an HTTP response writer.
func WriteResponseBody(write http.ResponseWriter, response interface{}) {

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	IsPanicErrors(err)
}
