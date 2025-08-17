package database

import (
	"github.com/Paya-4970/Portfolio-v.3/config"
	"github.com/Paya-4970/Portfolio-v.3/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB(cfg config.Config) *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	// AutoMigrate مدل‌ها
	if err := DB.AutoMigrate(
		&models.Profile{},
		&models.Shop{},
		&models.Product{},
	); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
	return DB
}
