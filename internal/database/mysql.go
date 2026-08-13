package database

import (
	"database/sql"
	"log"
	"projetoinfiel/internal/database/queries"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type service struct {
	db *sqlx.DB
}

func NewMySQLConnection() *sqlx.DB {
	dbPath := "sqlite.db"

	connection := sqlx.MustConnect("sqlite", dbPath)

	dbInstance := &service{
		db: connection,
	}

	dbInstance.createTable("Areas", queries.CreateAreaQuery)
	dbInstance.createTable("Users", queries.CreateUserTableQuery)
	dbInstance.createTable("Area_Users", queries.CreateAreaUserQuery)

	return connection
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) createTable(name, createTableQuery string) {
	_, err := s.db.Exec(createTableQuery)
	if err != nil {
		log.Println("Nao foi possivel criar a tabela: " + name)
		panic(err)
	}
}

func (s *service) Exec(query string, param ...any) (sql.Result, error) {
	return s.db.Exec(query, param...)
}

func (s *service) Query(query string, args ...any) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}
