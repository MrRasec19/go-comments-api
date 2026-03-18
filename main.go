package main

import (
	"database/sql"
	"fmt"
	"go-comments-api/internal/service"
	"go-comments-api/internal/store"
	"go-comments-api/internal/transport"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect to SQLITE
	db, err := sql.Open("sqlite3", "./comments.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Cret table "users" if no exist
	q := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		lastName TEXT NOT NULL,
		userName TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER,
		status INTEGER,
		created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Deoendencies inyection
	usersStore := store.New(db)
	usersService := service.New(usersStore)
	usersHandler := transport.New(usersService)

	// Routs setting
	http.HandleFunc("/users", usersHandler.HandleUsers)
	http.HandleFunc("/users/", usersHandler.HandleUserByID)

	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	fmt.Println("API Endpoints:")
	fmt.Println("	GET	/users		- get all users")
	fmt.Println("	POST	/users		- Create a new user")
	fmt.Println("	GET	/users/{id}	- get an specific users by id")
	fmt.Println("	PUT	/users/{id}	- Update a user")
	fmt.Println("	DELETE	/users/{id}	- Delete a user by id")

	// Start listening server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
