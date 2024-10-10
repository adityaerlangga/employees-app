package controllers

import (
    "employees-app/models"
    "employees-app/services"
    "employees-app/utils"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type EmployeeController struct {
    service services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
    return &EmployeeController{service: service}
}

func (c *EmployeeController) GetAllEmployees(ctx *gin.Context) {
    employees, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	if len(employees) == 0 {
		utils.RespondSuccess(ctx, http.StatusOK, "No employees found", nil)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Employees retrieved successfully", employees)	
}

func (c *EmployeeController) GetEmployeeByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid ID")
        return
    }

    employee, err := c.service.GetByID(id)
    if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, "Failed to retrieve employee")
        return
    }

	utils.RespondSuccess(ctx, http.StatusOK, "Employee retrieved successfully", employee)
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
    var employee models.Employee
    if err := ctx.ShouldBindJSON(&employee); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid request body")
        return
    }

	employee, err := c.service.Create(employee)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondSuccess(ctx, http.StatusCreated, "Employee created successfully", employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid ID")
        return
    }

    var employee models.Employee
    if err := ctx.ShouldBindJSON(&employee); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, err.Error())
        return
    }

	employee, err = c.service.Update(id, employee)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Employee updated successfully", employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid ID")
        return
    }

    if err := c.service.Delete(id); err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, err.Error())
        return
    }

	utils.RespondSuccess(ctx, http.StatusOK, "Employee deleted successfully", nil)
}
