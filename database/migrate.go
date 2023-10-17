package database

import (
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}

// func Rollback() {
// 	// Implementasikan rollback migrasi jika diperlukan
// 	DB.Migrator().DropTable(&models.User{}, &models.Photo{})
// }
