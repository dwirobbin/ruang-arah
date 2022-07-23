package conn

import (
	"database/sql"
	"fmt"
	"log"
	"ruang-arah/backend/config"
	"ruang-arah/backend/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dataSourceName(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, dbName,
	)
}

func createDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open(config.DB_DRIVER, dataSourceName(""))
	helper.PanicIfError(err)

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err = db.Exec(query)
	helper.PanicIfError(err)

	defer db.Close()

	return sql.Open(config.DB_DRIVER, dataSourceName(dbName))
}

func Connect() (*sql.DB, error) {
	db, err := createDB(config.DB_NAME)
	helper.PanicIfError(err)

	log.Printf("Successfully created database: %s.db", config.DB_NAME)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Println("Successfully connected to database")

	return db, nil
}
