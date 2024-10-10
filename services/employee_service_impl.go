package services

import (
    "employees-app/models"
    "gorm.io/gorm"
)

type EmployeeServiceImpl struct {
    db *gorm.DB
}

func (s *EmployeeServiceImpl) GetAll() ([]models.Employee, error) {
    var employees []models.Employee
    result := s.db.Find(&employees)
    return employees, result.Error
}

func (s *EmployeeServiceImpl) GetByID(id int) (models.Employee, error) {
    var employee models.Employee
    result := s.db.First(&employee, id)
    return employee, result.Error
}

func (s *EmployeeServiceImpl) Create(employee models.Employee) (models.Employee, error) {
    result := s.db.Create(&employee)
    return employee, result.Error
}

func (s *EmployeeServiceImpl) Update(id int, employee models.Employee) (models.Employee, error) {
    var existing models.Employee
    if err := s.db.First(&existing, id).Error; err != nil {
        return models.Employee{}, err
    }

    existing.Name = employee.Name
    existing.Position = employee.Position
    existing.Salary = employee.Salary

    result := s.db.Save(&existing)
    return existing, result.Error
}

func (s *EmployeeServiceImpl) Delete(id int) error {
    var employee models.Employee
    if err := s.db.First(&employee, id).Error; err != nil {
        return err
    }

    result := s.db.Delete(&employee)
    return result.Error
}
