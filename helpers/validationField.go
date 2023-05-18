package helpers

import (
	"AutomobilesAPI/dtos/request"
	"fmt"
)

// The function validates a request to create an automobile and returns an error if any required fields
// are missing or have invalid values.

func ValidateAutosCreateRequest(request *request.AutosCreateRequest) error {
	var emptyFields []string

	if request.Brand == "" || len(request.Brand) == 0 {
		emptyFields = append(emptyFields, "The field Brand is required.")
	}

	if request.Model == "" || len(request.Model) == 0 {
		emptyFields = append(emptyFields, "The field Model is required.")
	}

	if request.Price == 0 || request.Price <= 0 {
		emptyFields = append(emptyFields, "The field Price is required.")
	}
	if request.Year == 0 {
		emptyFields = append(emptyFields, "The field Year is required.")
	}

	if request.Year <= 1990 {
		emptyFields = append(emptyFields, "The pk must be greater than  1990. ")
	}

	if len(emptyFields) > 0 {
		return fmt.Errorf("The following fields are empty: %v", emptyFields)
	}

	return nil
}

// The function validates an AutosUpdateRequest object and returns an error if any required fields are
// empty or if the year is less than 1990.
func ValidateAutosUpdateRequest(request *request.AutosUpdateRequest) error {
	var emptyFields []string

	if request.AutoID == 0 || request.AutoID <= 0 {
		emptyFields = append(emptyFields, "The field AutoID is required.")
	}
	if request.Brand == "" || len(request.Brand) == 0 {
		emptyFields = append(emptyFields, "The field Brand is required.")
	}

	if request.Model == "" || len(request.Model) == 0 {
		emptyFields = append(emptyFields, "The field Model is required.")
	}

	if request.Price == 0 || request.Price <= 0 {
		emptyFields = append(emptyFields, "The field Price is required.")
	}

	if request.Year == 0 {
		emptyFields = append(emptyFields, "The field Year is required.")
	}

	if request.Year >= 1990 {
		emptyFields = append(emptyFields, "The pk must be greater than  1990. ")
	}

	if len(emptyFields) > 0 {
		return fmt.Errorf("The following fields are empty: %v", emptyFields)
	}

	return nil
}
