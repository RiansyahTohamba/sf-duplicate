package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	seedUser(db, "rian", "rian@gmail.com")
	seedUser(db, "bio", "bio@gmail.com")
	seedUser(db, "pogba", "pogba@gmail.com")

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		fmt.Println(rows)
	}
}

func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/sfduplicate.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			role_id INT   			
		);
		
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func seedUser(db *sql.DB, name, email string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO user (username, email, password, role_id) VALUES (?, ?, ?, 1)", name, email, hashedPassword)
	if err != nil {
		panic(err)
	}

}
