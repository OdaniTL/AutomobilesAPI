package repository

import (
	"AutomobilesAPI/models"
	"context"
)

var (
	querySelectById    = "SELECT AutoID,Brand, Model, Year, BodyType, EngineType, Displacement, Price,State FROM Autos WHERE  AutoID= $1"
	querySelectByBrand = `SELECT AutoID,Brand, Model, Year, BodyType, EngineType, Displacement, Price,State FROM Autos WHERE  State=true AND brand LIKE '%'||$1||'%'`
	querySelectAll     = "SELECT AutoID,Brand, Model, Year, BodyType, EngineType, Displacement, Price,State FROM Autos WHERE State=true"
	queryInsert        = "INSERT INTO Autos (Brand, Model, Year, BodyType, EngineType, Displacement, Price,State) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING AutoID"
	queryValidateId    = "SELECT  AutoID FROM Autos WHERE  AutoID= $1"
	queryUpdate        = "UPDATE Autos SET Brand=$1, Model=$2, Year=$3, BodyType=$4, EngineType=$5, Displacement=$6, Price=$7, state=$8 WHERE AutoID= $9"
	queryDelete        = "DELETE FROM Autos WHERE AutoID=$1"
)

// IAutosRepository is an interface that has methods Create, Update, FindById, FindByName, GetAll and
// Delete.
// @property {error} Create - This method is used to create a new auto.
// @property {error} Update - This method is used to update the auto details.
// @property FindById - This method will be used to find a auto by its ID.
// @property FindByName - This method will be used to find a auto by its name.
// @property GetAll - This method will return all the auto in the database.
// @property {error} Delete - This method is used to delete a auto from the database.

type IAutosRepository interface {
	Create(ctx context.Context, car models.Autos) error
	Update(ctx context.Context, car models.Autos) error
	FindById(ctx context.Context, autoId int) (models.Autos, error)
	FindByName(ctx context.Context, brandName string) ([]models.Autos, error)
	GetAll(ctx context.Context) ([]models.Autos, error)
	Delete(ctx context.Context, car int) error
}
