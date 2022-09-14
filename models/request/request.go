package request

type InsertCustomers struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
}

type UpdateCustomers struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
}
