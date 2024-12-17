package model

type OrderStatus string

type Order struct {
	ID            string        `gorm:"primaryKey" json:"id"`
	Status        OrderStatus   `json:"status"`
	ProductOrders []ProdukOrder `json:"product_orders"`
	ReferenceID   string        `gorm:"unique" json:"reference_id"`
}

type ProdukOrderStatus string

type ProdukOrder struct {
	ID         string            `gorm:"primaryKey" json:"id"`
	OrderID    string            `json:"order_id"`
	OrderCode  string            `json:"order_code"`
	Quantity   int               `json:"quantity"`
	TotalPrice int64             `json:"total_price"`
	Status     ProdukOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	OrderID string `json:"order_id"`
}
