package repository

import (
	"context"
	"database/sql"
	"log"
	"restful-api2/database"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

type customerClient struct {
	DB *sql.DB
}

type Repository interface {
	Insert(ctx context.Context, req request.InsertCustomers) (*response.InsertResponse, error)
	Update(ctx context.Context, req request.UpdateCustomers) (*response.UpdateResponse, error)
	Delete(ctx context.Context, customerId string) error
	FindById(ctx context.Context, customerId string) (*response.GetByIDResponse, error)
	FindAll(ctx context.Context) ([]response.GetAllResponse, error)
}

func NewClient() *customerClient {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &customerClient{
		DB: db,
	}
}
