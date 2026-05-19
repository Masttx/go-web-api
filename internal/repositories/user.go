package repositories

import (
	"database/sql"
	"fmt"
	"projetoinfiel/internal/database"
)

type userRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(name string, email string) (sql.Result, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	result, err := r.db.Exec(query, name, email)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}

	return result, nil
}

func (r *userRepository) List() ([]interface{}, error) {
	query := "SELECT * FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}

	var users []interface{}
	for rows.Next() {
		var user interface{}
		users = append(users, user)
	}

	return users, nil
}

// func Update()
// func Delete()
