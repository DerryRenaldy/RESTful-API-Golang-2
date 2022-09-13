package repository

import (
	"context"
	"database/sql"
	"restful-api2/helper"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Insert(ctx context.Context, req request.InsertCustomers) (response.InsertResponse, error) {
	query := "INSERT INTO restful_api.customers (customerId, phoneNumber, status) VALUES (?,?,?)"
	_, err := r.DB.ExecContext(ctx, query, req.CustomerID, req.PhoneNumber, req.Status)
	helper.PrintError(err)

	//response := response.InsertResponse{
	//	CustomerID:  req.CustomerID,
	//	PhoneNumber: req.PhoneNumber,
	//	Status:      req.Status,
	//}

}

func (r *Repository) Update(ctx context.Context, req request.UpdateCustomers) (response.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, customerId int) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) FindById(ctx context.Context, customerId int) response.GetByIDResponse {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) FindAll(ctx context.Context) []response.GetAllResponse {
	//TODO implement me
	panic("implement me")
}
