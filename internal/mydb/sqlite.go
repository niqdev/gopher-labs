package mydb

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go

type User struct {
	ID       int
	Username string
	Email    string
}

func SQLiteCrud() {
	dbPath := "./local/db/mysqlite.db"

	os.MkdirAll("./local/db/", 0755)
	os.Remove(dbPath)

	fmt.Println("connecting to SQLite:", dbPath)
	db, err := connect(dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := initSchema(db); err != nil {
		log.Fatalln(err)
	}

	if err := createUser(db, "gopher", "gopher@example.com"); err != nil {
		log.Fatalln(err)
	}

	users, err := readUsers(db)
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		fmt.Println(fmt.Sprintf("[%d] Username=%s Email=%s", user.ID, user.Username, user.Email))
	}

	if err := updateUser(db, 1, "updated_user", "updated@example.com"); err != nil {
		log.Fatalln(err)
	}

	if err := deleteUser(db, 1); err != nil {
		log.Fatalln(err)
	}
}

func connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening db")
	}

	// ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrapf(err, "error connecting to db")
	}

	return db, nil
}

func initSchema(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL
    );
    `
	_, err := db.Exec(query)
	return errors.Wrapf(err, "error init schema")
}

func createUser(db *sql.DB, username, email string) error {
	_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	return errors.Wrapf(err, "error creating user")
}

func readUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.Wrapf(err, "error reading users")
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, errors.Wrapf(err, "error scanning row")
		}
		users = append(users, user)
	}

	return users, nil
}

func updateUser(db *sql.DB, id int, newUsername, newEmail string) error {
	_, err := db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", newUsername, newEmail, id)
	return errors.Wrapf(err, "error updating user")
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return errors.Wrapf(err, "error deleting user")
}
