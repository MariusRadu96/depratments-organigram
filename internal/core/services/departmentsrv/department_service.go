package departmentservice

import (
	"context"
	"departments-organigram/internal/core/domain"
	"departments-organigram/internal/core/ports"
	"strings"
)

type departmentSrv struct {
	departmentsStore ports.DepartmentsStore
}

func NewDepartmentSrv(departmentsStore ports.DepartmentsStore) *departmentSrv {
	return &departmentSrv{
		departmentsStore: departmentsStore,
	}
}

func (d *departmentSrv) CreateDepartment(ctx context.Context, department domain.Department) error {
	return d.departmentsStore.CreateDepartment(ctx, department)
}

func (d *departmentSrv) UpdateDepartment(ctx context.Context, department domain.Department) error {
	return d.departmentsStore.UpdateDepartment(ctx, department)
}

func (d *departmentSrv) DeleteDepartment(ctx context.Context, id int) error {
	return d.departmentsStore.DeleteDepartment(ctx, id)
}

func (d *departmentSrv) GetDepartmentByID(ctx context.Context, id int) (domain.Department, error) {

	departmentsHierarchy, err := d.departmentsStore.GetDepartmentHierarchy(ctx, id)
	if err != nil {
		return domain.Department{}, err
	}

	departmentHierarchySlice := []string{}

	for _, department := range departmentsHierarchy {
		departmentHierarchySlice = append(departmentHierarchySlice, department.Name)
	}

	department, err := d.departmentsStore.GetDepartmentByID(ctx, id)
	if err != nil {
		return domain.Department{}, err
	}

	department.Hierarchhy = strings.Join(departmentHierarchySlice, " --- ")
	return department, nil
}

func (d *departmentSrv) GetAllDepartments(ctx context.Context) ([]domain.Department, error) {
	return d.departmentsStore.GetAllDepartments(ctx)
}
