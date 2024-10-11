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
$ docker-compose up -d
```

| No | Task                              | Check             |
|:--:|----------------------------------|:-----------------:|
| 1  | Database Setup (Postgre SQL)    |        ✅         |
| 2  | API Implementation               |        ✅         |
| 3  | Dockerize the Application        |        ✅         |
| 4  | Deploy to Kubernetes             | Not Implemented   |
| 5  | Version Control with Github      |        ✅         |
| 6  | Bonus (Optional): Integrate Cloud Infrastructure |        ✅         |


### Routes:

https://www.postman.com/atom-pos/iykra-technical-test-aditya-erlangga-wibowo/overview

- URL Local: http://localhost:8080
- URL Cloud: https://iykra.adityaerlangga.my.id

  
    

|                          METHOD                          |              URL               |   Description   |
| :------------------------------------------------------: | :----------------------------: | :-------------: |
|      ![GET](https://img.shields.io/badge/-GET-blue)      |  https://iykra.adityaerlangga.my.id/employees   |   GetAllEmployees   |
|      ![GET](https://img.shields.io/badge/-GET-blue)      | https://iykra.adityaerlangga.my.id/employees/:id |   GetEmployeeByID   |
|      ![POST](https://img.shields.io/badge/-POST-success)      | https://iykra.adityaerlangga.my.id/employees/ |   CreateEmployee 
|     ![PUT](https://img.shields.io/badge/-PUT-orange)     |   https://iykra.adityaerlangga.my.id/employees/:id   | UpdateEmployee |
![DELETE](https://img.shields.io/badge/-DELETE-critical) | https://iykra.adityaerlangga.my.id/employees/:id  | DeleteEmployee  |

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

## Infra Spesifications

- Digital Ocean VPS RAM 2GB/1vCPU
- Ubuntu 20.04
- Domain Propagation
- Redirect HTTPS