package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/db_btpns?parseTime=true"
	// Gantilah "username", "password", dan "database_name" sesuai dengan kredensial MySQL Anda.

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Auto-migrate the models
	// DB.AutoMigrate(&models.User{}, &models.Photo{})
}
