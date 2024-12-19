package main

import (
	"crypto/rand"
	"crypto/rsa"
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	mRepo "go-restaurant-app/internal/repository/menu"
	oRepo "go-restaurant-app/internal/repository/order"
	uRepo "go-restaurant-app/internal/repository/user"
	rUsecase "go-restaurant-app/internal/usecase/resto"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {
	db := database.GetDB(dbAddress)
	// sebaiknya ditempatkan di .inv
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
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
