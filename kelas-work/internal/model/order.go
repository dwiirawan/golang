package model

type OrderStatus string

type Order struct {
	ID            string        `gorm:"primaryKey" json:"id"`
	Status        OrderStatus   `json:"status"`
	ProductOrders []ProdukOrder `json:"product_orders"`
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
	OrderCode string `json:"order_status"`
	Quantity  int    `json:"quntity"`
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
}

type GetOrderInfoRequest struct {
	OrderID string `json:"order_id"`
}
