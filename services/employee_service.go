package services

import (
    "employees-app/models"
    "gorm.io/gorm"
)

type EmployeeService interface {
    GetAll() ([]models.Employee, error)
    GetByID(id int) (models.Employee, error)
    Create(employee models.Employee) (models.Employee, error)
    Update(id int, employee models.Employee) (models.Employee, error)
    Delete(id int) error
}

func NewEmployeeService(db *gorm.DB) EmployeeService {
    return &EmployeeServiceImpl{db: db}
}
