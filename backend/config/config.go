package config

import "time"

const (
	// APP CONFIGURATION
	APP_NAME = "Ruang-Arah"
	API_PORT = "5000"

	// JWT CONFIGURATION
	SECRET_KEY    = "secret-key"
	TOKEN_EXPIRES = time.Minute * 60

	// DATABASE CONFIGURATION
	DB_DRIVER = "mysql"
	DB_USER   = "root"
	DB_PASS   = "rootinit"
	DB_HOST   = "localhost"
	DB_PORT   = "3306"
	DB_NAME   = "ruang_arah_v2"
)
