package main

import (
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	mRepo "go-restaurant-app/internal/repository/menu"
	oRepo "go-restaurant-app/internal/repository/order"
	uRepo "go-restaurant-app/internal/repository/user"
	rUsecase "go-restaurant-app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {
	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32)
	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	e := echo.New()
	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
