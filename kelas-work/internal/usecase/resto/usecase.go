package resto

import (
	"go-restaurant-app/internal/model"
)

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	Order(request model.GetOrderMenuRequest) (model.Order, error)
	GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error)
}
