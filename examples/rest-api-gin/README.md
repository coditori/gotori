# Golang Restful API implementation
Here I used below technologies:
- Gin framework
- Swagger (Swaggo)
- SQLite (in-memory)
- Testify

## How to run
Before running you can download dependencies by:
```console
foo@bar:~$ go get .
```
Then generate swagger docs by

```console
foo@bar:~$ swag init
```
Then simply run the application

```console
foo@bar:~$ go run main.go
```

## Default user info
username/password: admin
You can also login with this user once and enter the token inside swagger/postman then you have the access to all endpoints

## Check if the API is running
Check this on browser: http://localhost:8000/ping

## Swagger UI
It can be accessed by: http://localhost:8000/docs/index.html

## How to test
Here I used actual end-to-end test for all possible cases, it's a good idea to add mock test later and mock the repository layer
```console
foo@bar:~$ go test ./...
```
