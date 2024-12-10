package menu

import "go-restaurant-app/internal/model"

type Repository interface {
	GetMenuList(MenuType string) ([]model.MenuItem, error)
	GetMenu(OrderCode string) (model.MenuItem, error)
}
