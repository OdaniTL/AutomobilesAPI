package services

import (
	"AutomobilesAPI/dtos/request"
	"AutomobilesAPI/dtos/response"
	"context"
)

// IAutoServices is an interface with 6 methods.
// @property {error} Create - Create a new auto
// @property {error} Udate - This is the method that will be used to update the auto.
// @property FindById - Finds a auto by its ID.
// @property FindByName - Finds a auto by brand name.
// @property GetAll - This method is used to get all the autos from the database.
// @property {error} Delete - Delete the auto by ID
type IAutoServices interface {
	Create(ctx context.Context, request request.AutosCreateRequest) error
	Udate(ctx context.Context, request request.AutosUpdateRequest) error
	FindById(ctx context.Context, autoId int) (response.AutosResponse, error)
	FindByName(ctx context.Context, brandName string) ([]response.AutosResponse, error)
	GetAll(ctx context.Context) ([]response.AutosResponse, error)
	Delete(ctx context.Context, autoId int) error
}
