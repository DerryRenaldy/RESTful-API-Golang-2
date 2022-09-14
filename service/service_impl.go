package service

import (
	"context"
	"restful-api2/helper"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

func (s *Service) Create(ctx context.Context, req request.InsertCustomers) (*response.InsertResponse, error) {
	status, err := s.Repository.Insert(ctx, req)
	helper.PrintError(err)

	status.CustomerID = req.CustomerID
	status.PhoneNumber = req.PhoneNumber

	return status, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateCustomers) (*response.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, customerId string) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindById(ctx context.Context, customerId string) *response.GetByIDResponse {
	response, err := s.Repository.FindById(ctx, customerId)
	helper.PrintError(err)

	return response
}

func (s *Service) FindAll(ctx context.Context) []response.GetAllResponse {
	//TODO implement me
	panic("implement me")
}
