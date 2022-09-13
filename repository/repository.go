package repository

import (
	"context"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

type RepositoryInter interface {
	Insert(ctx context.Context, req request.InsertCustomers) (response.InsertResponse, error)
	Update(ctx context.Context, req request.UpdateCustomers) (response.UpdateResponse, error)
	Delete(ctx context.Context, customerId int)
	FindById(ctx context.Context, customerId int) response.GetByIDResponse
	FindAll(ctx context.Context) []response.GetAllResponse
}
