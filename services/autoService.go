package services

import (
	"AutomobilesAPI/dtos/request"
	"AutomobilesAPI/dtos/response"
	"AutomobilesAPI/helpers"
	"AutomobilesAPI/models"
	"AutomobilesAPI/repository"
	"context"
	"errors"
	"strings"
)

type AutoServices struct {
	AutoReporsitory repository.IAutosRepository
}

func NewAutoServices(autoRepository repository.IAutosRepository) IAutoServicesMock {
	return &AutoServices{AutoReporsitory: autoRepository}
}

// Create implements IAutoService
func (p *AutoServices) Create(ctx context.Context, request request.AutosCreateRequest) error {

	isValid := request.Validate()
	if isValid != nil {

		return isValid
	}
	car := models.Autos{
		Brand:        request.Brand,
		Model:        request.Model,
		Year:         request.Year,
		BodyType:     request.EngineType,
		EngineType:   request.EngineType,
		Displacement: request.Displacement,
		Price:        request.Price,
		State:        request.State,
	}
	return p.AutoReporsitory.Create(ctx, car)
}

// Delete implements IAutoService
// This function is implementing the `Delete` method of the `IAutoService` interface. It takes a
// context and an integer `autoId` as input parameters and returns an error. It calls the `Delete`
// method of the `AutoRepository` to delete the automobile with the given `autoId` from the database.
func (p *AutoServices) Delete(ctx context.Context, autoId int) error {
	if autoId == 0 || autoId <= 0 {
		return errors.New("the field autoid is required")
	}
	return p.AutoReporsitory.Delete(ctx, autoId)
}

// FindById implements IAutoService
func (p *AutoServices) FindById(ctx context.Context, autoId int) (response.AutosResponse, error) {

	if autoId == 0 || autoId <= 0 {
		return response.AutosResponse{}, errors.New("the field autoid is required")
	}
	resp, err := p.AutoReporsitory.FindById(ctx, autoId)

	if err != nil {
		return response.AutosResponse{}, err
	}

	return helpers.MappingDataResponse(resp), nil
}

// FindByName implements IAutoService
func (p *AutoServices) FindByName(ctx context.Context, brandName string) ([]response.AutosResponse, error) {

	if len(brandName) == 0 || len(strings.TrimSpace(brandName)) == 0 {
		return []response.AutosResponse{}, errors.New("the field brand is required")
	}
	resutlData, err := p.AutoReporsitory.FindByName(ctx, brandName)
	if err != nil {
		return []response.AutosResponse{}, err
	}

	return helpers.MappingDataResponses(resutlData), nil
}

// GetAll implements IAutoService
func (p *AutoServices) GetAll(ctx context.Context) ([]response.AutosResponse, error) {
	resutlData, err := p.AutoReporsitory.GetAll(ctx)
	if err != nil {
		return []response.AutosResponse{}, err
	}
	return helpers.MappingDataResponses(resutlData), nil
}

// Udate implements IAutoService
func (p *AutoServices) Udate(ctx context.Context, request request.AutosUpdateRequest) error {

	isValid := request.Validate()

	if isValid != nil {

		return isValid
	}

	car := models.Autos{
		AutoID:       request.AutoID,
		Brand:        request.Brand,
		Model:        request.Model,
		Year:         request.Year,
		BodyType:     request.EngineType,
		EngineType:   request.EngineType,
		Displacement: request.Displacement,
		Price:        request.Price,
		State:        request.State,
	}
	return p.AutoReporsitory.Update(ctx, car)
}
