package repositories

import (
	"database/sql"
	"fmt"
	"projetoinfiel/internal/types"
)

type AreaUserRepository struct {
	db *sql.DB
}

func NewAreaUserRepository(db *sql.DB) *AreaUserRepository {
	return &AreaUserRepository{
		db: db,
	}
}

func (r *AreaUserRepository) Create(areaID int64, userID int64) (sql.Result, error) {
	query := "INSERT INTO area_users (area_id, user_id) VALUES (?, ?)"

	result, err := r.db.Exec(query, areaID, userID)
	if err != nil {
		return nil, fmt.Errorf("Error to insert area_user: %v", err)
	}

	return result, nil
}

func (r *AreaUserRepository) ListUsersByArea(areaID int64) ([]types.User, error) {
	query := `
		SELECT u.id, u.name, u.email 
		FROM users u 
		INNER JOIN area_users au ON u.id = au.user_id 
		WHERE au.area_id = ?`

	rows, err := r.db.Query(query, areaID)
	if err != nil {
		return nil, fmt.Errorf("Error to list users by area: %v", err)
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("Error scanning user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *AreaUserRepository) ListAreasByUser(userID int64) ([]types.Area, error) {
	query := `
		SELECT a.id, a.name, a.description 
		FROM areas a 
		INNER JOIN area_users au ON a.id = au.area_id 
		WHERE au.user_id = ?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Error to list areas by user: %v", err)
	}
	defer rows.Close()

	var areas []types.Area
	for rows.Next() {
		var area types.Area
		err = rows.Scan(&area.ID, &area.Name, &area.Description)
		if err != nil {
			return nil, fmt.Errorf("Error scanning area: %v", err)
		}
		areas = append(areas, area)
	}

	return areas, nil
}

func (r *AreaUserRepository) Delete(areaID int64, userID int64) error {

	query := "DELETE FROM area_users WHERE area_id = ? AND user_id = ?"

	res, err := r.db.Exec(query, areaID, userID)
	if err != nil {
		return fmt.Errorf("Error to delete area_user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Nenhuma relação encontrada entre a área %d e o usuário %d para deletar", areaID, userID)
	}

	return nil
}
