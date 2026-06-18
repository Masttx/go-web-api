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
	createAreaQuery := `CREATE TABLE IF NOT EXISTS areas (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,description TEXT)`
	s.createTable("Areas", createAreaQuery)

	createUserQuery := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,
	email TEXT NOT NULL)`
	s.createTable("Users", createUserQuery)

	createAreaUserQuery := `CREATE TABLE IF NOT EXISTS area_users (area_id INTEGER REFERENCES area(id), user_id INTEGER REFERENCES user(id)`
	s.createTable("Area_Users", createAreaUserQuery)
}

func (s *service) createTable(name, createTableQuery string) {
	_, err := s.db.Exec(createTableQuery)
	if err != nil {
		log.Println("Nao foi possivel criar a tabela: " + name)
	}
}

func (s *service) Exec(query string, param ...any) (sql.Result, error) {
	return s.db.Exec(query, param...)
}

func (s *service) Query(query string, args ...any) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}
