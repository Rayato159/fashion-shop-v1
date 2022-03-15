package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/configs"
)

func DBConnect(cfg *configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", cfg.Database.Url)
	if err != nil {
		return nil, fmt.Errorf("error, can't connect to database, %w", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, can't send ping to database, %w", err)
	}

	log.Println("database has connected successfully!")
	return db, nil
}
