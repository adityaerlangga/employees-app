package main

import (
    "employees-app/controllers"
    "employees-app/services"
    "employees-app/utils"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    db := utils.InitDB()

    employeeService := services.NewEmployeeService(db)
    employeeController := controllers.NewEmployeeController(employeeService)

    r.GET("/employees", employeeController.GetAllEmployees)
    r.GET("/employees/:id", employeeController.GetEmployeeByID)
    r.POST("/employees", employeeController.CreateEmployee)
    r.PUT("/employees/:id", employeeController.UpdateEmployee)
    r.DELETE("/employees/:id", employeeController.DeleteEmployee)

    // Run server
    r.Run(":8080")
}
