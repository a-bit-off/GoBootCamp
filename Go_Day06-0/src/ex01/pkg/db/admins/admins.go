package admins

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Admins"
)

type AdminData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AddNewAdmin(newAdmin AdminData) {
	db := connectToPostgreSQL()
	createUsersTable(db)
	insertUsersTable(db, newAdmin)
	closeDB(db)
}

func connectToPostgreSQL() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlConn)
	checkError(err)

	err = db.Ping()
	checkError(err)
	return db
}

func createUsersTable(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS Users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(50),
		password VARCHAR(50)
	)
	`
	_, err := db.Exec(createTableQuery)
	checkError(err)
}

func insertUsersTable(db *sql.DB, newAdmin AdminData) {
	insertDataQuery := `
		INSERT INTO Users (login, password) VALUES ($1, $2)
	`
	_, err := db.Exec(insertDataQuery, newAdmin.Name, newAdmin.Password)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func closeDB(db *sql.DB) {
	db.Close()
}
