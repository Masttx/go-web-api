package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"projetoinfiel/internal/database/queries"
	"projetoinfiel/internal/types"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name string, email string, password_hash string) (sql.Result, error) {
	result, err := r.db.Exec(queries.CreateUserQuery, name, email, password_hash)
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

func (r *UserRepository) ExistsUserByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, queries.ExistsUserByEmailQuery, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserRepository) FindUserByEmail(email string) (types.User, error) {
	var user types.User
	err := r.db.Get(&user, queries.FindUserByEmailQuery, email)
	if err != nil {
		return types.User{}, fmt.Errorf("Error to find user: %v", err)
	}
	return user, nil
}

func (r *UserRepository) List(ctx context.Context) ([]types.User, error) {
	var result []types.User
	err := r.db.SelectContext(ctx, &result, queries.ListUsersQuery)
	if err != nil {
		return nil, fmt.Errorf("Error to list users: %v", err)
	}

	return result, nil
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

func (r *UserRepository) UpdateUserArea(id int64, areaId int64) error {
	query := "UPDATE users SET area_id = ? WHERE id = ?"

	_, err := r.db.Exec(query, areaId, id)
	if err != nil {
		return fmt.Errorf("Error to update user: %v", err)
	}

	return nil
}

func (r *UserRepository) ListByArea(areaId int64) ([]types.User, error) {
	query := "SELECT * FROM users WHERE area_id = ?"

	rows, err := r.db.Query(query, areaId)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}
	defer rows.Close()

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
