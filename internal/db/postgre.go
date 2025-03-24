package db

import (
	"fmt"

	"log"

	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/models/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	AutoMigrate(dst ...interface{}) error
	GetDB() *gorm.DB
}

type GormDatabase struct {
	Conn *gorm.DB
}

func (g *GormDatabase) AutoMigrate(dst ...interface{}) error {
	return g.Conn.AutoMigrate(dst...)
}

func (g *GormDatabase) GetDB() *gorm.DB {
	return g.Conn
}

func ConnectDatabase(cfg config.Config) Database {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=%s", cfg.Database.Host, cfg.Database.User, cfg.Database.Name, cfg.Database.Port, cfg.Database.Password, cfg.Database.SSLMode)

	conn, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	db := &GormDatabase{Conn: conn}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
