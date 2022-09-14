package repository

import (
	"context"
	"database/sql"
	"restful-api2/helper"
	"restful-api2/models/request"
	"restful-api2/models/response"
)

func (r *customerClient) Insert(ctx context.Context, req request.InsertCustomers) (*response.InsertResponse, error) {
	queryInsert := `INSERT INTO restful_api.customers (customerId, phoneNumber, status) VALUES (?,?,?)`
	query, err := r.DB.PrepareContext(ctx, queryInsert)
	helper.PrintError(err)

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	_, err = query.ExecContext(ctx, req.CustomerID, req.PhoneNumber, req.Status)
	helper.PrintError(err)

	res := &response.InsertResponse{}
	queryGetStatus := `SELECT description FROM restful_api.customers_status where status=?`
	err = r.DB.QueryRowContext(ctx, queryGetStatus, req.Status).Scan(&res.Status)
	helper.PrintError(err)

	return res, nil
}

func (r *customerClient) Update(ctx context.Context, req request.UpdateCustomers) (*response.UpdateResponse, error) {
	queryUpdate := `UPDATE restful_api.customers 
					SET phoneNumber = ?, status=?
					WHERE customerId=?;`

	query, err := r.DB.PrepareContext(ctx, queryUpdate)
	helper.PrintError(err)

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	_, err = query.ExecContext(ctx, req.PhoneNumber, req.Status, req.CustomerID)
	helper.PrintError(err)

	querySearch := `SELECT c.customerId, c.phoneNumber, cs.description 
					FROM restful_api.customers c 
					INNER JOIN restful_api.customers_status cs 
					ON c.status=cs.status WHERE c.customerId =?;`

	res := &response.UpdateResponse{}

	err = r.DB.QueryRowContext(ctx, querySearch, req.CustomerID).Scan(&res.CustomerID, &res.PhoneNumber, &res.Status)
	helper.PrintError(err)

	return res, nil
}

func (r *customerClient) Delete(ctx context.Context, customerId string) error {
	queryDelete := `DELETE FROM restful_api.customers WHERE customerId= ?;`
	_, err := r.DB.ExecContext(ctx, queryDelete, customerId)
	helper.PrintError(err)

	return nil
}

func (r *customerClient) FindById(ctx context.Context, customerId string) (*response.GetByIDResponse, error) {
	queryFindId := `SELECT c.customerId, c.phoneNumber, cs.description 
					FROM restful_api.customers c 
					INNER JOIN restful_api.customers_status cs 
					ON c.status=cs.status WHERE c.customerId =?;`

	res := &response.GetByIDResponse{}
	err := r.DB.QueryRowContext(ctx, queryFindId, customerId).Scan(&res.CustomerID, &res.PhoneNumber, &res.Status)
	helper.PrintError(err)

	return res, nil
}

func (r *customerClient) FindAll(ctx context.Context) ([]response.GetAllResponse, error) {
	queryFindAll := `SELECT c.customerId, c.phoneNumber, cs.description 
					FROM restful_api.customers c 
					INNER JOIN restful_api.customers_status cs 
					ON c.status=cs.status;`
	rows, err := r.DB.QueryContext(ctx, queryFindAll)
	helper.PrintError(err)

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var resSlice []response.GetAllResponse

	for rows.Next() {
		res := response.GetAllResponse{}
		err := rows.Scan(&res.CustomerID, &res.PhoneNumber, &res.Status)
		helper.PrintError(err)
		resSlice = append(resSlice, res)
	}

	return resSlice, nil
}
