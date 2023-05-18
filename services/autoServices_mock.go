package services

import (
	"AutomobilesAPI/dtos/request"
	"AutomobilesAPI/dtos/response"
	"AutomobilesAPI/repository"
	"context"
)

type IAutoServicesMock interface {
	Create(ctx context.Context, request request.AutosCreateRequest) error
	Udate(ctx context.Context, request request.AutosUpdateRequest) error
	FindById(ctx context.Context, autoId int) (response.AutosResponse, error)
	FindByName(ctx context.Context, brandName string) ([]response.AutosResponse, error)
	GetAll(ctx context.Context) ([]response.AutosResponse, error)
	Delete(ctx context.Context, autoId int) error
}

type AutoServicesMock struct {
	AutoReporsitory repository.RepositoryMocked
}

func NewAutoServicesMock(autoRepository repository.RepositoryMocked) IAutoServicesMock {
	return &AutoServicesMock{AutoReporsitory: autoRepository}
}

// Create implements IAutoServicesMock
func (*AutoServicesMock) Create(ctx context.Context, request request.AutosCreateRequest) error {
	return nil
}

// Delete implements IAutoServicesMock
func (*AutoServicesMock) Delete(ctx context.Context, autoId int) error {
	return nil
}

// FindById implements IAutoServicesMock
func (*AutoServicesMock) FindById(ctx context.Context, autoId int) (response.AutosResponse, error) {
	return response.AutosResponse{}, nil
}

// FindByName implements IAutoServicesMock
func (*AutoServicesMock) FindByName(ctx context.Context, brandName string) ([]response.AutosResponse, error) {
	return []response.AutosResponse{}, nil
}

// GetAll implements IAutoServicesMock
func (*AutoServicesMock) GetAll(ctx context.Context) ([]response.AutosResponse, error) {
	return []response.AutosResponse{}, nil
}

// Udate implements IAutoServicesMock
func (*AutoServicesMock) Udate(ctx context.Context, request request.AutosUpdateRequest) error {
	return nil
}
