package database

import (
	"fmt"
	"waysbean/models"
	"waysbean/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Profile{},
		&models.Transaction{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Gagal")
	}

	fmt.Println(("Migration Berhasil"))
}
