package main

import (
	"database/sql"
	"os"
)

func open() (*mysql, error) {
	database := os.Getenv("MYSQL_DATABASE")
	password := os.Getenv("MYSQL_PASSWORD")

	dataSourceName := "root:" + password + "@tcp(35.240.177.194:3306)/" + database
	dataSourceName += "?charset=utf8mb4"
	dataSourceName += "&collation=utf8mb4_unicode_ci"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &mysql{db: db}, nil
}
