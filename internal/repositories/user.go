package repositories

import (
	"database/sql"
	"fmt"
	"projetoinfiel/internal/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name string, email string) (sql.Result, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	result, err := r.db.Exec(query, name, email)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}

	return result, nil
}

func (r *UserRepository) Read(id int64) (types.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	var user types.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return types.User{}, fmt.Errorf("Error to find user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) List() ([]types.User, error) {
	query := "SELECT * FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}

	var users []types.User
	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("Error to scan user: %v", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Update(id int64, name string, email string) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"

	_, err := r.db.Exec(query, name, email, id)
	if err != nil {
		return fmt.Errorf("Error to update user: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = ?"

	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error to delete user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Nenhuma linha deletada. Usuário com ID %d não encontrado", id)
	}

	return nil
}
