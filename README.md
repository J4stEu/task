# Task

Simple task manager in Golang based on mocks. 

## Why

This project was made to brainstorm and show some common features:
 - Clean architecture based on layers and interfaces.
 - Models relations.
 - Mocked data. This pattern can be used within the first iteration of the project, especially when the database is not relevant.
 - REST HTTP API, JWT authentication with cookies management.
 - Some background work with contexts and goroutines.

## Layers

Architecture: Clean Architecture.

Project structure: [Golang Project Layout](https://github.com/golang-standards/project-layout).

`./internal/repository` -> layer to work with data. Mock realisation can be changed to Postgres etc.

`./internal/service` -> business logic layer.

`./internal/api` -> representation layer. 

## Run

```sh
make run
```

## Authenticate

Use `./internal/repository/token/token_mock.go`, change token value in line 25.

Where to get value? After using the login HTTP endpoint, you will get token value:

```sh
curl --location 'http://127.0.0.1:8085/api/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "login": "admin",
    "password": "test"
}'
```

Response example:
```json
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6IjE0ZjY1YTY2LTQ1NWQtNDRmMS1hMmJjLWYyY2Q4OTI5M2Q4NCJ9.zhMRiuqkloje2EuRQmJrqqVt-h_mIChfti63od6yGW0",
  "token": {
    ...
    "value": "14f65a66-455d-44f1-a2bc-f2cd89293d84",
    ...
  }
}
```

