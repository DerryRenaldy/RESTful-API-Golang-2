package service

import (
	"context"
	"restful-api2/helper"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

func (s *Service) Create(ctx context.Context, req request.InsertCustomers) (*response.InsertResponse, error) {
	resp, err := s.Repository.Insert(ctx, req)
	helper.PrintError(err)

	return resp, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateCustomers) (*response.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, customerId int) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindByIn(ctx context.Context, customerId int) *response.GetByIDResponse {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindAll(ctx context.Context) []response.GetAllResponse {
	//TODO implement me
	panic("implement me")
}
