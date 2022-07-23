package migrate

import (
	"database/sql"
	"log"
	"ruang-arah/backend/helper"
)

func Migrate(db *sql.DB) {
	queries := []string{
		`SET FOREIGN_KEY_CHECKS=0;`,
		`CREATE TABLE IF NOT EXISTS users (
				id INT NOT NULL AUTO_INCREMENT,
				username VARCHAR(255) NOT NULL,
				email VARCHAR(150) NOT NULL UNIQUE,
				password TEXT NOT NULL,
				role VARCHAR(15) NOT NULL,
				loggedin BOOLEAN NOT NULL DEFAULT FALSE,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE programming_languages (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(255) NOT NULL UNIQUE,
				image_url TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS questions (
				id INT NOT NULL AUTO_INCREMENT,
				proglang_id INT NOT NULL,
				question TEXT NOT NULL,
				correct_answer TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (proglang_id) REFERENCES programming_languages(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS incorrect_answers (
				id INT NOT NULL AUTO_INCREMENT,
				question_id INT NOT NULL,
				option_a TEXT NOT NULL,
				option_b TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS answers_attempts (
				id INT NOT NULL AUTO_INCREMENT,
				answer TEXT NOT NULL,
				question_id INT NOT NULL,
				user_id INT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS levels (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(15) NOT NULL UNIQUE,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS recommendations (
				id INT NOT NULL AUTO_INCREMENT,
			  image_url TEXT NOT NULL,
				level_id INTEGER NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (level_id) REFERENCES levels(id),
				PRIMARY KEY (id)
			);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		helper.PanicIfError(err)
	}

	log.Println("Successfully migrated all table")
}
