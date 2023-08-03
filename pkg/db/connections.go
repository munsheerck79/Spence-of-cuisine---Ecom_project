package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/munsheerck79/Ecom_project.git/pkg/config"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB variables
var (
	DB  *gorm.DB
	err error
)

// To connect database
func ConnToDB(cfg config.Config) (*gorm.DB, error) {

	dsn := cfg.DATABASE

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Failed to Connect Database")
		return nil, errors.New("failed to connect database")
	}
	fmt.Println("Successfully Connected to database")

	// Migrate models
	err := DB.AutoMigrate(
		// Users
		domain.Users{},
		domain.Admin{},
		domain.Category{},
		domain.Variation{},
		domain.Product{},
		domain.Price{},
		domain.Address{},
		domain.Cart{},
		domain.WishList{},
		domain.Coupon{},
		domain.OrderStatus{},
		domain.Orders{},
		domain.Wallet{},
		domain.WalletHistory{},
		domain.OrdersItems{},
		domain.OrdersTemp{},
		domain.OrdersItemsTemp{},

	)
	if err != nil {
		log.Fatal("Migration failed")
		return nil, nil
	}
	fmt.Println("DB migration success")
	return DB, nil

}
