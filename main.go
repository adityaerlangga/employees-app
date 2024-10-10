package main

import (
    "employees-app/controllers"
    "employees-app/services"
    "employees-app/utils"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
)

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://iykra.adityaerlangga.my.id"}, // Atur origin yang diizinkan
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

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
