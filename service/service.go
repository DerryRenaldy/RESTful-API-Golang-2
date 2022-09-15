package service

import (
	"context"
	"restful-api2/models/request"
	"restful-api2/models/response"
	"restful-api2/repository"
)

type Service struct {
	Repository repository.Repository
}

type Repository interface {
	Create(ctx context.Context, req request.InsertCustomers) (*response.InsertResponse, error)
	Update(ctx context.Context, req request.UpdateCustomers) (*response.UpdateResponse, error)
	Delete(ctx context.Context, customerId string) error
	FindById(ctx context.Context, customerId string) *response.GetByIDResponse
	FindAll(ctx context.Context) []response.GetAllResponse
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
