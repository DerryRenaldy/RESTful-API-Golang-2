package controller

import (
	"net/http"
	"restful-api2/service"
)

type Controller struct {
	Services service.ServiceInter
}

func NewController(services service.ServiceInter) *Controller {
	return &Controller{Services: services}
}

type ControllerInter interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}
