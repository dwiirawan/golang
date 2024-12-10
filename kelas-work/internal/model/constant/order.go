package constant

import "go-restaurant-app/internal/model"

const (
	OrderStatusProcessed model.OrderStatus = "processed"
	OrderStatusFinished  model.OrderStatus = "finished"
	OrderStatusFailed    model.OrderStatus = "failed"
)

const (
	ProdukOrderStatusPreparing model.ProdukOrderStatus = "preparing"
	ProdukOrderStatusFinished  model.ProdukOrderStatus = "finished"
)
