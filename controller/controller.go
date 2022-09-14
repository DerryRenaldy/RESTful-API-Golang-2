package controller

import (
	"net/http"
	"restful-api2/service"
)

type Controller struct {
	Services service.Repository
}

type Repository interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}

func NewController(services service.Repository) *Controller {
	return &Controller{Services: services}
}
