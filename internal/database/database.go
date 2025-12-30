package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"go-setup/internal/config"

	_ "github.com/lib/pq"
)

func NewPostgres(cfg *config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	log.Println("Database connected")
	return db, nil
}