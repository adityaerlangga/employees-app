# Go - Employee App

![Logo](./assets/banner.png)


## Setup
```
$ git clone https://github.com/adityaerlangga/employees-app
$ cd employees-app
```

```
DB_HOST=db # FOR RUNNING DOCKER
# DB_HOST=localhost # FOR LOCAL ENV
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=db_employees
DB_PORT=5432
```

```
$ docker-compose up --build
```

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
    "position": "Software Engineer",
    "salary": 20000
}
```

##### UPDATE (PARAMS ID)

```
{
    "name": "Aditya",
    "position": "Full Stack Engineer",
    "salary": 15000000
}
```
