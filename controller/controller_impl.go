package controller

import (
	"encoding/json"
	"net/http"
	"restful-api2/helper"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	customerCreateRequest := request.InsertCustomers{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(customerCreateRequest)
	helper.PrintError(err)

	customerResponse, err := c.Services.Create(r.Context(), customerCreateRequest)
	webResponse := response.InsertResponse{
		CustomerID:  customerResponse.CustomerID,
		PhoneNumber: customerResponse.PhoneNumber,
		Status:      customerResponse.Status,
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PrintError(err)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) FindById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
