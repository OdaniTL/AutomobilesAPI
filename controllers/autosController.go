package controllers

import (
	"AutomobilesAPI/dtos/request"
	"AutomobilesAPI/dtos/response"
	"AutomobilesAPI/helpers"
	"AutomobilesAPI/services"
	"net/http"
	"strconv"
)

type AutosController struct {
	autoServices services.IAutoServices
}

func NewAutosController(autoServs services.IAutoServices) *AutosController {
	return &AutosController{autoServices: autoServs}
}

func (contller *AutosController) Create(w http.ResponseWriter, r *http.Request) {

	if !helpers.IsMethodsAllowed(r.Method, http.MethodPost, w) {
		return
	}
	autoRequest := request.AutosCreateRequest{}
	helpers.ReadRequestBody(r, &autoRequest)

	apiResponse := response.OkRequest("Ok")
	err := contller.autoServices.Create(r.Context(), autoRequest)
	if err != nil {
		apiResponse = response.GetBadRequest(err.Error())
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	helpers.WriteResponseBody(w, apiResponse)
}

func (contller *AutosController) Update(w http.ResponseWriter, r *http.Request) {

	if !helpers.IsMethodsAllowed(r.Method, http.MethodPut, w) {
		return
	}
	apiResponse := response.ApiResponse{}
	autoRequet := request.AutosUpdateRequest{}
	helpers.ReadRequestBody(r, &autoRequet)
	err := contller.autoServices.Udate(r.Context(), autoRequet)
	if err != nil {
		apiResponse = response.GetBadRequest(err.Error())
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	helpers.WriteResponseBody(w, response.OkRequest("OK"))
}

func (contller *AutosController) Delete(w http.ResponseWriter, r *http.Request) {

	if !helpers.IsMethodsAllowed(r.Method, http.MethodDelete, w) {
		return
	}
	_Id := r.URL.Query().Get("autoId")
	autoId, err := strconv.Atoi(_Id)
	apiResponse := response.ApiResponse{}

	if err != nil {
		apiResponse = response.GetBadRequest(err.Error())
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	apiResponse = response.OkRequest("Ok")

	err = contller.autoServices.Delete(r.Context(), autoId)
	if err != nil {
		apiResponse = response.NotFoundRequest("Auto id not found.")
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	helpers.WriteResponseBody(w, apiResponse)
}

func (contller *AutosController) FindById(w http.ResponseWriter, r *http.Request) {

	if !helpers.IsMethodsAllowed(r.Method, http.MethodGet, w) {
		return
	}
	apiResponse := response.ApiResponse{}
	_Id := r.URL.Query().Get("autoid")
	autoID, err := strconv.Atoi(_Id)
	if err != nil {
		apiResponse = response.GetBadRequest("Erro to convert string to int.")
		helpers.WriteResponseBody(w, apiResponse)
		return
	}

	result, err := contller.autoServices.FindById(r.Context(), autoID)
	if err != nil {
		apiResponse = response.NotFoundRequest("Auto id not found.")
		//apiResponse.Status = "Not Found"
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	apiResponse = response.OkRequest("Ok")
	apiResponse.Data = result
	helpers.WriteResponseBody(w, apiResponse)
}

func (contller *AutosController) FindByName(w http.ResponseWriter, r *http.Request) {

	if !helpers.IsMethodsAllowed(r.Method, http.MethodGet, w) {
		return
	}

	brandName := r.URL.Query().Get("name")

	apiResponse := response.OkRequest("Ok")
	if brandName == "" || len(brandName) == 0 {
		apiResponse = response.GetBadRequest("The name field is requerit")
		apiResponse.Code = http.StatusBadRequest
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	result, err := contller.autoServices.FindByName(r.Context(), brandName)

	if err != nil {
		apiResponse = response.NotFoundRequest("brand Name not found.")
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	apiResponse.Data = result
	helpers.WriteResponseBody(w, apiResponse)
}

func (contller *AutosController) GetAll(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsMethodsAllowed(r.Method, http.MethodGet, w) {
		return
	}
	apiResponse := response.OkRequest("Ok")

	result, err := contller.autoServices.GetAll(r.Context())

	if err != nil {
		apiResponse = response.NotFoundRequest("Auto not found.")
		helpers.WriteResponseBody(w, apiResponse)
		return
	}
	apiResponse.Data = result
	helpers.WriteResponseBody(w, apiResponse)
}
