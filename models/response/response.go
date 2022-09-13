package response

type InsertResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type GetAllResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type GetByIDResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type UpdateResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
