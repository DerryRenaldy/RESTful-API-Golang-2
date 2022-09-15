package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"restful-api2/controller"
	"restful-api2/helper"
	"restful-api2/repository"
	"restful-api2/service"
)

type Server struct {
	sql        repository.Repository
	services   service.Repository
	controller controller.Repository
}

func Register() *Server {
	SVR := &Server{}

	SVR.sql = repository.NewClient()

	SVR.services = service.NewService(SVR.sql)

	SVR.controller = controller.NewController(SVR.services)

	return SVR
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/insertCustomer", s.controller.Create).Methods(http.MethodPost)
	r.HandleFunc("/getByIdCustomer/{customerId}", s.controller.FindById).Methods(http.MethodGet)
	r.HandleFunc("/getAll", s.controller.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/updateCustomer/{customerId}", s.controller.Update).Methods(http.MethodPut)
	r.HandleFunc("/deleteCustomer/{customerId}", s.controller.Delete).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8010", r)
	helper.PrintError(err)
}
