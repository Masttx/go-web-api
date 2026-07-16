package repositories

import (
	"database/sql"
	"fmt"
	"projetoinfiel/internal/types"
)

type AreaRepository struct {
	db *sql.DB
}

func NewAreaRepository(db *sql.DB) *AreaRepository {
	return &AreaRepository{
		db: db,
	}
}

func (r *AreaRepository) Create(name, description string) (sql.Result, error) {
	query := "INSERT INTO areas (name, description) VALUES (?, ?)"

	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, fmt.Errorf("Error to insert area: %v", err)
	}

	return result, nil
}

func (r *AreaRepository) Read(id int64) (types.Area, error) {
	query := "SELECT * FROM areas WHERE id = ?"

	var area types.Area
	err := r.db.QueryRow(query, id).Scan(&area.ID, &area.Name, &area.Description)
	if err != nil {
		return types.Area{}, fmt.Errorf("Error to find area: %v", err)
	}

	return area, nil
}

func (r *AreaRepository) List() ([]types.Area, error) {
	query := "SELECT * FROM areas"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error to insert areas: %v", err)
	}
	defer rows.Close()

	var areas []types.Area
	for rows.Next() {
		var area types.Area
		err = rows.Scan(&area.ID, &area.Name, &area.Description)
		if err != nil {
			return nil, fmt.Errorf("Error to scan area: %v", err)
		}

		areas = append(areas, area)
	}

	return areas, nil
}

func (r *AreaRepository) Update(id int64, name, description string) error {
	query := "UPDATE areas SET name = ?, description = ? WHERE id = ?"

	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return fmt.Errorf("Error to update area: %v", err)
	}

	return nil
}

func (r *AreaRepository) Delete(id int64) error {
	query := "DELETE FROM areas WHERE id = ?"

	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error to delete area: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Nenhuma linha deletada. Area com ID %d não encontrado", id)
	}

	return nil
}
