package services

import (
    "employees-app/models" // Import models package
    "gorm.io/gorm" // Import gorm package
)

// EmployeeService defines the interface for employee operations
type EmployeeService interface {
    GetAll() ([]models.Employee, error)
    GetByID(id int) (models.Employee, error)
    Create(employee models.Employee) (models.Employee, error)
    Update(id int, employee models.Employee) (models.Employee, error)
    Delete(id int) error
}

// NewEmployeeService creates a new instance of EmployeeService
func NewEmployeeService(db *gorm.DB) EmployeeService {
    return &EmployeeServiceImpl{db: db}
}
