package ports

import (
	"context"
	"departments-organigram/internal/core/domain"
)

type DepartmentService interface {
	CreateDepartment(ctx context.Context, department domain.Department) error
	UpdateDepartment(ctx context.Context, department domain.Department) error
	DeleteDepartment(ctx context.Context, id int) error
	GetDepartmentByID(ctx context.Context, id int) (domain.Department, error)
	GetAllDepartments(ctx context.Context) ([]domain.Department, error)
}
