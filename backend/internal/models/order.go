package models

import (
	"time"

	"github.com/google/uuid"
)

type MpesaSTKPushRequest struct {
	OrderID uuid.UUID `json:"order_id"`
	Amount  float64   `json:"amount"`
	Phone   string    `json:"phone"`
}

type MpesaSTKPushResponse struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	MpesaRef      string    `json:"mpesa_ref"`
	Status        string    `json:"status"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type Order struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	UserName    string      `json:"user_name"`
	Status      string      `json:"status"`
	TotalAmount float64     `json:"total_amount"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Product   *Product  `json:"product,omitempty"`
}

type OrderRequest struct {
	Items []struct {
		ProductID uuid.UUID `json:"product_id"`
		Quantity  int       `json:"quantity"`
		UnitPrice float64   `json:"unit_price"`
	} `json:"items"`
	ShippingAddress string `json:"shipping_address"`
	PhoneNumber     string `json:"phone_number"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}
