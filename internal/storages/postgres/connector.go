package postgres

import (
	"database/sql"
	"fmt"
	"gw_exchanger/pkg/env"
	"log"
	"os"
)

type Connector struct {
	USERNAME_DB string
	PASSWORD_DB string
	DATABASE    string
	HOST_DB     string
	PORT_DB     string
}

func (c *Connector) Connection() (*sql.DB, error) {
	err := env.LoadConfig("../config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	conInfo := Connector{
		USERNAME_DB: os.Getenv("USERNAME_DB"),
		PASSWORD_DB: os.Getenv("PASSWORD_DB"),
		DATABASE:    os.Getenv("DATABASE"),
		HOST_DB:     os.Getenv("HOST_DB"),
		PORT_DB:     os.Getenv("PORT_DB"),
	}
	conString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		conInfo.HOST_DB, conInfo.PORT_DB, conInfo.USERNAME_DB,
		conInfo.PASSWORD_DB, conInfo.DATABASE,
	)
	db, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, fmt.Errorf("Incorrect data: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
