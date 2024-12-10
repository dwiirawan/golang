package main

import (
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	mRepo "go-restaurant-app/internal/repository/menu"
	oRepo "go-restaurant-app/internal/repository/order"
	rUsecase "go-restaurant-app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {
	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)

	e := echo.New()
	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
