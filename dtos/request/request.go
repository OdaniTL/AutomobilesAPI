package request

import (
	"fmt"
	"strconv"
)

type FieldValidator interface {
	Validate() error
}

type Field struct {
	Name  string
	Value *string
}

func (f *Field) Validate() error {
	if f.Value == nil || *f.Value == "" {
		return fmt.Errorf("el campo '%s' está vacío. \n", f.Name)
	}
	if *f.Value == "0" {
		return fmt.Errorf("el campo '%s' está vacío.\n", f.Name)
	}
	if *f.Value == "c0" {

		return fmt.Errorf("The field %s  is required.\n", f.Name)
	}
	if *f.Value == "y0" {

		return fmt.Errorf("The field %s must be greater than  1990.\n", f.Name)
	}
	return nil
}

func ValidateFields(fields []FieldValidator) error {
	var errs []error
	for _, field := range fields {
		if err := field.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

/*
	if request.Year == 0 {
			emptyFields = append(emptyFields, "The field Year is required.")
		}

		if request.Year >= 1990 {
			emptyFields = append(emptyFields, "The pk must be greater than  1990. ")
		}
*/
func (req *AutosCreateRequest) Validate() error {
	fields := []FieldValidator{
		&Field{Name: "Brand", Value: &req.Brand},
		&Field{Name: "Model", Value: &req.Model},
		&Field{Name: "BodyType", Value: &req.BodyType},
		&Field{Name: "EngineType", Value: &req.EngineType},
		&Field{Name: "Price", Value: func() *string {
			price := strconv.FormatFloat(float64(req.Price), 'f', -1, 32)
			return &price
		}()},
		&Field{Name: "Year", Value: CustomValidate(req.Year)},
		&Field{Name: "Year", Value: CustomValidateYear(req.Year)},
	}
	return ValidateFields(fields)
}

func (req *AutosUpdateRequest) Validate() error {
	fields := []FieldValidator{
		&Field{Name: "Brand", Value: &req.Brand},
		&Field{Name: "Model", Value: &req.Model},
		&Field{Name: "BodyType", Value: &req.BodyType},
		&Field{Name: "EngineType", Value: &req.EngineType},
		&Field{Name: "Price", Value: func() *string {
			price := strconv.FormatFloat(float64(req.Price), 'f', -1, 32)
			return &price
		}()},
		&Field{Name: "AutoID", Value: CustomValidate(req.AutoID)},
		&Field{Name: "Year", Value: CustomValidate(req.Year)},
	}
	return ValidateFields(fields)
}
func CustomValidate(fieldVuale int) *string {

	emptyFields := "check"
	if fieldVuale == 0 || fieldVuale <= 0 {
		emptyFields = "c0"
	}

	return &emptyFields
}
func CustomValidateYear(fieldVuale int) *string {

	emptyFields := "check"
	if fieldVuale > 1000 && fieldVuale <= 1990 {
		emptyFields = "y0"
	}

	return &emptyFields
}
