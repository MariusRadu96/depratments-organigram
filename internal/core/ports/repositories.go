package ports

import (
	"context"
	"departments-organigram/internal/core/domain"
)

type DepartmentsStore interface {
	CreateDepartment(ctx context.Context, department domain.Department) error
	UpdateDepartment(ctx context.Context, department domain.Department) error
	DeleteDepartment(ctx context.Context, departmentID int) error
	GetDepartmentByID(ctx context.Context, id int) (domain.Department, error)
	GetDepartmentHierarchy(ctx context.Context, departmentID int) ([]domain.Department, error)
	GetAllDepartments(ctx context.Context) ([]domain.Department, error)
}

type UsersStore interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}
