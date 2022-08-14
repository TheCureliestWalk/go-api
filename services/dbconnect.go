package services

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, *sql.DB, error) {
	host := "127.0.0.1"
	user := "postgres"
	pass := "12345678"
	name := "postgres"
	port := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("Error connecting to DB -> %+v\n", err.Error())
		return nil, nil, err
	} else {
		sqlDB, _ := db.DB()
		return db, sqlDB, nil
	}
}
