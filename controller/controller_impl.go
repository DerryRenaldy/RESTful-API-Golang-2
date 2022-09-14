package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restful-api2/helper"
	"restful-api2/models/request"
)

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	customerCreateRequest := request.InsertCustomers{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerCreateRequest)
	helper.PrintError(err)

	customerResponse, err := c.Services.Create(r.Context(), customerCreateRequest)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(customerResponse)
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
	params := mux.Vars(r)
	customerId := params["customerId"]

	//custIdRequest := request.GetById{
	//	CustomerID: customerId,
	//}

	customerResponse := c.Services.FindById(r.Context(), customerId)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(customerResponse)
	helper.PrintError(err)
}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
