package departments

import (
	"context"
	"database/sql"
	"departments-organigram/internal/core/domain"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type store struct {
	db *sql.DB
}

func NewStore(database *sql.DB) *store {
	return &store{
		db: database,
	}
}

func (s *store) CreateDepartment(ctx context.Context, department domain.Department) error {
	_, err := s.db.Query("CALL CreateDepartment(?, ?, ?)", department.Name, department.ParentID, department.Flags)
	return err
}

func (s *store) UpdateDepartment(ctx context.Context, department domain.Department) error {
	_, err := s.db.Query("CALL UpdateDepartment(?, ?, ?, ?)", department.ID, department.Name, department.ParentID, department.Flags)
	return err
}

func (s *store) DeleteDepartment(ctx context.Context, id int) error {
	_, err := s.db.Query("CALL DeleteDepartment(?)", id)
	return err
}

func (s *store) GetDepartmentByID(ctx context.Context, id int) (domain.Department, error) {
	rows, err := s.db.Query("CALL GetDepartmentByID(?)", id)
	if err != nil {
		return domain.Department{}, err
	}
	defer rows.Close()

	var department domain.Department
	for rows.Next() {
		if err := rows.Scan(
			&department.ID,
			&department.Name,
			&department.ParentID,
			&department.Flags,
			&department.CreatedAt,
			&department.UpdatedAt); err != nil {
			return domain.Department{}, err
		}
	}

	if department.ID == 0 {
		return domain.Department{}, errors.New("not found")
	}

	return department, nil
}

func (r *store) GetDepartmentHierarchy(ctx context.Context, departmentID int) ([]domain.Department, error) {
	rows, err := r.db.Query("CALL GetDepartmentHierarchy(?)", departmentID)
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return nil, err
	}
	defer rows.Close()

	var departments []domain.Department

	for rows.Next() {
		var department domain.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
			&department.ParentID,
			&department.Flags,
			&department.CreatedAt,
			&department.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		departments = append(departments, department)

	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return departments, nil
}

func (r *store) GetAllDepartments(ctx context.Context) ([]domain.Department, error) {
	rows, err := r.db.Query("CALL GetAllDepartments()")
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return nil, err
	}
	defer rows.Close()

	var departments []domain.Department

	for rows.Next() {
		var department domain.Department

		err := rows.Scan(
			&department.ID,
			&department.Name,
			&department.ParentID,
			&department.Flags,
			&department.CreatedAt,
			&department.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		departments = append(departments, department)

	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return departments, nil
}
