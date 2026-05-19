package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type service struct {
	db *sql.DB
}

func NewMySQLConnection() *sql.DB {
	dbPath := "sqlite.db"

	connection, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	dbInstance := &service{
		db: connection,
	}

	dbInstance.initDatabase()

	return connection
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) initDatabase() {
	createUserQuery := "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,email TEXT NOT NULL)"

	_, err := s.db.Exec(createUserQuery)
	if err != nil {
		log.Println("Nao foi possivel criar a tabela")
	}
}

func (s *service) Exec(query string, param ...any) (sql.Result, error) {
	return s.db.Exec(query, param...)
}

func (s *service) Query(query string, args ...any) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}
