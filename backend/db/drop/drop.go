package drop

import (
	"database/sql"
	"log"
	"ruang-arah/backend/helper"
)

func Drop(db *sql.DB) {
	queries := []string{
		`SET FOREIGN_KEY_CHECKS = 0;`,
		`DROP TABLE IF EXISTS users;`,
		`DROP TABLE IF EXISTS programming_languages;`,
		`DROP TABLE IF EXISTS questions;`,
		`DROP TABLE IF EXISTS incorrect_answers;`,
		`DROP TABLE IF EXISTS answers_attempts;`,
		`DROP TABLE IF EXISTS levels;`,
		`DROP TABLE IF EXISTS recommendations;`,
		`SET FOREIGN_KEY_CHECKS = 1;`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		helper.PanicIfError(err)
	}

	log.Println("Successfully dropped all table")
}
