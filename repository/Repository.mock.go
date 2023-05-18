package repository

import (
	"AutomobilesAPI/models"
	"context"
)

type RepositoryMocked struct {
}

func (rp *RepositoryMocked) Create(ctx context.Context, car models.Autos) error {
	return nil
}
func (rp *RepositoryMocked) Update(ctx context.Context, car models.Autos) error {
	return nil
}
func (rp *RepositoryMocked) FindById(ctx context.Context, productId int) (models.Autos, error) {
	return models.Autos{}, nil
}
func (rp *RepositoryMocked) FindByName(ctx context.Context, productName string) ([]models.Autos, error) {
	return []models.Autos{}, nil
}
func (rp *RepositoryMocked) GetAll(ctx context.Context) ([]models.Autos, error) {
	return []models.Autos{}, nil
}
func (rp *RepositoryMocked) Delete(ctx context.Context, car int) error {
	return nil
}
