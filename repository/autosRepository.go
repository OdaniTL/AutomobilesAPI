package repository

import (
	"AutomobilesAPI/helpers"
	"AutomobilesAPI/models"
	"context"
	"database/sql"
	"errors"
)

type AutosRepository struct {
	DB *sql.DB
}

func NewAutosRepository(Db *sql.DB) IAutosRepository {
	return &AutosRepository{DB: Db}
}

// This function is implementing the `FindById` method of the `IAutosRepository` interface. It takes a
// context and an integer `autoID` as input parameters and returns a `models.Autos` struct and an
// error. It retrieves the automobile from the database that matches the `autoID` provided and returns
// it as a `models.Autos` struct. If the automobile with the provided `autoID` does not exist in the
// database, it returns an error.
func (p *AutosRepository) FindById(ctx context.Context, autoID int) (models.Autos, error) {
	tx, err := p.DB.Begin()

	var autos models.Autos
	if err != nil {
		return autos, err
	}

	defer helpers.CommitOrRollback(tx)
	result, err := tx.QueryContext(ctx, querySelectById, autoID)

	if err != nil {
		return autos, err
	}

	defer result.Close()
	if result.Next() {

		err = result.Scan(
			&autos.AutoID,
			&autos.Brand,
			&autos.Model,
			&autos.Year,
			&autos.BodyType,
			&autos.EngineType,
			&autos.Displacement,
			&autos.Price,
			&autos.State)

		if err != nil {
			return autos, err
		}

		return autos, nil
	} else {

		return autos, errors.New("auto id not found")
	}

}

// The `FindByName` function is implementing the `FindByName` method of the `IAutosRepository`
// interface. It takes a context and a string `brandName` as input parameters and returns a slice of
// `models.Autos` struct and an error. It retrieves all the automobiles from the database that match
// the `brandName` provided and returns them as a slice of `models.Autos` struct. If there is an error
// during the retrieval process, it returns an error.
func (p *AutosRepository) FindByName(ctx context.Context, brandName string) ([]models.Autos, error) {

	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer helpers.CommitOrRollback(tx)

	result, errr := tx.QueryContext(ctx, querySelectByBrand, brandName)
	var autos []models.Autos

	if errr != nil {
		return autos, errr
	}
	defer result.Close()
	for result.Next() {
		car := models.Autos{}
		errr = result.Scan(
			&car.AutoID,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.BodyType,
			&car.EngineType,
			&car.Displacement,
			&car.Price,
			&car.State)

		if errr != nil {
			return autos, errr
		}
		autos = append(autos, car)
	}
	return autos, nil
}

// The `GetAll` function is implementing the `GetAll` method of the `IAutosRepository` interface. It
// takes a context as input parameter and returns a slice of `models.Autos` struct and an error. It
// retrieves all the automobiles from the database and returns them as a slice of `models.Autos`
// struct. If there is an error during the retrieval process, it returns an error.
func (p *AutosRepository) GetAll(ctx context.Context) ([]models.Autos, error) {
	tx, err := p.DB.Begin()
	var autos []models.Autos

	if err != nil {
		return autos, err
	}
	defer helpers.CommitOrRollback(tx)

	result, errr := tx.QueryContext(ctx, querySelectAll)

	if errr != nil {
		return autos, errr
	}
	defer result.Close()
	for result.Next() {
		car := models.Autos{}
		err = result.Scan(
			&car.AutoID,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.BodyType,
			&car.EngineType,
			&car.Displacement,
			&car.Price,
			&car.State)
		if err != nil {
			return autos, err
		}
		autos = append(autos, car)
	}
	return autos, nil
}

// The `Create` function is implementing the `Create` method of the `IAutosRepository` interface. It
// takes a context and a `models.Autos` struct as input parameters and returns an error. It inserts a
// new automobile into the database with the information provided in the `models.Autos` struct. It
// returns an error if the operation is not completed correctly.
func (p *AutosRepository) Create(ctx context.Context, car models.Autos) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	defer helpers.CommitOrRollback(tx)
	lastInsertId := 0

	queryResult := tx.QueryRowContext(
		ctx,
		queryInsert,
		car.Brand,
		car.Model,
		car.Year,
		car.BodyType,
		car.EngineType,
		car.Displacement,
		car.Price,
		car.State,
	).Scan(&lastInsertId)

	if queryResult != nil {
		return queryResult
	}

	if lastInsertId == 0 {
		return errors.New("The update operation has not been completed correctly")
	}

	return nil
}

// The `Update` function is implementing the `Update` method of the `IAutosRepository` interface. It
// takes a context and a `models.Autos` struct as input parameters and returns an error. It updates the
// information of an existing automobile in the database by checking if the `AutoID` exists in the
// database, and if it does, it updates the information of the automobile with the new information
// provided in the `models.Autos` struct. If the `AutoID` does not exist in the database, it returns an
// error.
func (p *AutosRepository) Update(ctx context.Context, car models.Autos) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollback(tx)
	var exist bool

	if err := tx.QueryRowContext(ctx, queryValidateId, car.AutoID).Scan(&exist); err != nil {
		return errors.New("Autoid not found")
	}
	result, errr := tx.ExecContext(
		ctx,
		queryUpdate,
		car.Brand,
		car.Model,
		car.Year,
		car.BodyType,
		car.EngineType,
		car.Displacement,
		car.Price,
		car.State,
		car.AutoID)
	if errr != nil {
		return errr
	}

	rwsAffect, err := result.RowsAffected()

	if err != nil {
		return err
	}
	if rwsAffect == 0 {
		return errors.New("Autos id not found")
	}

	return nil
}

// This function is implementing the `Delete` method of the `IAutosRepository` interface. It takes a
// context and an integer `autoID` as input parameters and returns an error.
func (p *AutosRepository) Delete(ctx context.Context, autoID int) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}
	defer helpers.CommitOrRollback(tx)
	result, err := tx.ExecContext(ctx, queryDelete, autoID)

	if err != nil {
		return err
	}
	result.RowsAffected()

	rwsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rwsAffect == 0 {
		return errors.New("Autos id not found")
	}
	return nil

}
