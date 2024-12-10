package model

type OrderStatus string

type Order struct {
	ID            string `gorm:"primaryKey"`
	Status        OrderStatus
	ProductOrders []ProdukOrder
}

type ProdukOrderStatus string

type ProdukOrder struct {
	ID         string `gorm:"primaryKey"`
	OrderID    string
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProdukOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode string
	Quantity  int
}

type GetOrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest
}

type GetOrderInfoRequest struct {
	OrderID string
}
