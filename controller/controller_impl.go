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
	customerUpdateRequest := request.UpdateCustomers{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerUpdateRequest)
	helper.PrintError(err)

	params := mux.Vars(r)
	customerId := params["customerId"]

	customerUpdateRequest.CustomerID = customerId

	customerResponse, err := c.Services.Update(r.Context(), customerUpdateRequest)
	helper.PrintError(err)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(customerResponse)
	helper.PrintError(err)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId := params["customerId"]

	err := c.Services.Delete(r.Context(), customerId)
	helper.PrintError(err)
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
	customerResponse := c.Services.FindAll(r.Context())
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(customerResponse)
	helper.PrintError(err)
}
