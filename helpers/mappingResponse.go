package helpers

import (
	"AutomobilesAPI/dtos/response"
	"AutomobilesAPI/models"
)

// The function maps data from a model of Autos to a response object of AutosResponse.
func MappingDataResponse(model models.Autos) response.AutosResponse {

	return response.AutosResponse{
		AutoID:       model.AutoID,
		Brand:        model.Brand,
		Model:        model.Model,
		Year:         model.Year,
		BodyType:     model.BodyType,
		EngineType:   model.EngineType,
		Displacement: model.Displacement,
		Price:        model.Price,
		State:        model.State,
	}
}

// The function maps data from a list of car models to a list of car response objects.
func MappingDataResponses(cars []models.Autos) []response.AutosResponse {

	var autos []response.AutosResponse
	if len(cars) == 0 {
		return autos
	}
	for _, auto := range cars {

		card := response.AutosResponse{
			AutoID:       auto.AutoID,
			Brand:        auto.Brand,
			Model:        auto.Model,
			Year:         auto.Year,
			BodyType:     auto.BodyType,
			EngineType:   auto.EngineType,
			Displacement: auto.Displacement,
			Price:        auto.Price,
			State:        auto.State,
		}
		autos = append(autos, card)
	}
	return autos
}
