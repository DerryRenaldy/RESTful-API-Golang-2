package repository

import (
	"context"
	"database/sql"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

type Repository struct {
	DB *sql.DB
}

type RepositoryInter interface {
	Insert(ctx context.Context, req request.InsertCustomers) (response.InsertResponse, error)
	Update(ctx context.Context, req request.UpdateCustomers) (response.UpdateResponse, error)
	Delete(ctx context.Context, customerId int) error
	FindById(ctx context.Context, customerId int) (response.GetByIDResponse, error)
	FindAll(ctx context.Context) ([]response.GetAllResponse, error)
}
