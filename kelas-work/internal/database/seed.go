package database

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	// Migrate to schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProdukOrder{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam rica-rica",
			OrderCode: "ayam_rica-rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Air mineral",
			OrderCode: "air_mineral",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
