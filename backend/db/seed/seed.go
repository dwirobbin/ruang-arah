package seed

import (
	"database/sql"
	"log"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/security"
)

func Seed(db *sql.DB) {
	query := `INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?);`

	hashed := security.GeneratePasswordHash("admin123")
	args := []interface{}{"Admin", "admin@gmail.com", hashed, "admin"}

	_, err := db.Exec(query, args...)
	helper.PanicIfError(err)

	queries := []string{
		`INSERT INTO programming_languages (name, image_url)
		VALUES
		('Go', 'https://www.linkpicture.com/q/go_1.png'),
		('Python', 'https://www.linkpicture.com/q/python_1.png'),
		('Java', 'https://www.linkpicture.com/q/java_3.png'),
		('C#', 'https://www.linkpicture.com/q/cs_2.png'),
		('Ruby', 'https://www.linkpicture.com/q/ruby_1.png'),
		('PHP', 'https://www.linkpicture.com/q/php_2.png'),
		('Rust', 'https://www.linkpicture.com/q/rust_5.png'),
		('JavaScript', 'https://www.linkpicture.com/q/js_16.png');`,

		`INSERT INTO questions (proglang_id, question, correct_answer)
		VALUES
		(1, 'Apa itu Go ?', "Go adalah bahasa pemrograman yang dibuat oleh Google"),
		(1, 'Yang bukan termasuk pointer ?', "#"),
		(1, 'Apa itu variable constant di Go ?', "Variable yang nilainya bersifat tetap"),
		(1, 'Manakah penulisan fungsi di Go yang benar ?', "func namaFungsi()"),
		(1, 'Di Go, berapa jumlah maksimal yang bisa direturn oleh sebuah fungsi ?', "Tak Terhingga"),
		(1, 'Fungsi mana dari paket runtime yang digunakan untuk mendapatkan jumlah prosesor yang digunakan dalam program ?', "fmt.Println(runtime.NumCPU())"),
		(1, 'Manakah di antara pernyataan berikut yang benar untuk Golang ?', "Golang does not provide support for method overloading and type inheritance"),
		(1, 'Manakah di antara opsi di bawah ini yang mencetak tipe data variabel a untuk kode di bawah ini ?', "fmt.Printf('Type of a is %T\n', a)"),
		(1, 'Apakah mungkin untuk mendeklarasikan variabel dari tipe yang berbeda dalam satu baris kode di Golang ?', "Ya, bisa"),
		(1, 'Perulangan yang tidak ada di bahasa Go, tetapi ada di bahasa selain Go ?', "while")`,

		`INSERT INTO incorrect_answers (question_id, option_a, option_b)
		VALUES
		(1, 'Go adalah bahasa pemrograman yang dibuat oleh Microsoft', 'Go adalah bahasa pemrograman yang dibuat oleh Apple'),
		(2, '*', '&'),
		(3, 'Variable yang nilainya bersifat dinamis', 'Variable yang nilainya bisa bersifat dinamis & tetap'),
		(4, 'function namaFungsi()', 'def namaFungsi()'),
		(5, '2', '5'),
		(6, 'fmt.Println(runtime.CPU())', 'fmt.Println(runtime.CPUCount())'),
		(7, 'Golang provides support for method overloading and type inheritance', 'Golang provides support for method overloading and does not support type inheritance'),
		(8, 'fmt.Printf("Type of a is %X\n", a)', 'fmt.Printf("Type of a is %D\n", a)'),
		(9, 'Tidak bisa', 'Mustahil'),
		(10, 'for', 'for range')`,

		`INSERT INTO levels (name)
		VALUES 
		('Beginner'), ('Intermediate');`,

		`INSERT INTO recommendations (image_url, level_id)
		VALUES
		('https://www.linkpicture.com/q/Basic_1.png', 1),
		('https://www.linkpicture.com/q/Intermediate.png', 2);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		helper.PanicIfError(err)
	}

	log.Println("Successfully seeded all table")
}
