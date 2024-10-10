# Go - Employee App

![Logo](./assets/banner.png)

### Routes

|                          METHOD                          |              URL               |   Description   |
| :------------------------------------------------------: | :----------------------------: | :-------------: |
|      ![GET](https://img.shields.io/badge/-GET-blue)      |  http://localhost:8080/employees   |   GetAllEmployees   |
|      ![GET](https://img.shields.io/badge/-GET-blue)      | http://localhost:8080/employees/:id |   GetEmployeeByID   |
|      ![POST](https://img.shields.io/badge/-POST-success)      | http://localhost:8080/employees/ |   CreateEmployee 
|     ![PUT](https://img.shields.io/badge/-PUT-orange)     |   http://localhost:8080/employees/:id   | UpdateEmployee |
![DELETE](https://img.shields.io/badge/-DELETE-critical) | http://localhost:8080/employees/:id  | DeleteEmployee  |

### PAYLOAD

##### CREATE

```
{
    "name": "Aditya",
    "position": "Erlangga",
    "salary": 20000
}
```

##### UPDATE

```
{
    "name": "Aditya",
    "position": "Full Stack Engineer",
    "salary": 15000000
}
```

## Requirements

- go 1.19.1

## Setup

Do this :

```
$ git clone https://github.com/adityaerlangga/go-clean-architecture
$ go mod download
$ go run main.go
```
