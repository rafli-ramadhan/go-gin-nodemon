# Golang + Gin + Gorm + Nodemon + PostgreSQL

Requirement :
- Go 1.16
- Node.js
- PostgreSQL

# Start 🚀

## Install Modules

```bash
go mod download && go mod tidy && go mod verify
```

If the message below was shown, do the next step.
```
go: finding module for package github.com/forkyid/go-rest-api/docs
github.com/forkyid/go-rest-api/src/route imports
        github.com/forkyid/go-rest-api/docs: no matching versions for query "latest"
```

## Swagger Installation and Swag Initialization

```bash
go install github.com/swaggo/swag/cmd/swag@v1.6.7
```

```bash
swag init -g src/main.go
```

```bash
go mod tidy
```

## Install Nodemon

```bash
npm install -g nodemon
```

## Running the Server

### Go run + Nodemon

```bash
nodemon --exec go run src/main.go --signal SIGTERM
```

![running](https://user-images.githubusercontent.com/112603532/221372094-f0c58450-b82b-425b-a49b-e9aab5f121a9.png)

Swagger API Documentation URL:
```url
http://localhost:5000/swagger/index.html#/
```

![swagger](https://user-images.githubusercontent.com/112603532/221371951-adc780d5-fc53-4b60-ac8c-7cadc412e484.png)

### Docker

```bash
docker-compose up --build
```

## Repository Structure

```bash
.
├── .github
│   └── PULL_REQUEST_TEMPLATE.md
├── src
│   ├── connection
│   │   └── connection.go
│   ├── constant
│   │   └── constant.go
│   ├── controller
│   │   └── v1
│   │       ├── user.go
│   ├── models
│   │   └── v1
│   │       ├── user.go
│   ├── pkg
│   │   └── http
│   │       ├── user.go
│   ├── routes
│   │   └── routes.go
│   ├── service
│   │   └── v1
│   │       ├── user.go
│   └── main.go
├── .env.example
├── .gitignore
├── README.md
└── go.mod
```
