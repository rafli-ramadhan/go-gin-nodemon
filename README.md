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

This command will create a new folder `docs`. Reference : [swaggo](https://github.com/swaggo/swag).

Then, run this command again:

```bash
go mod tidy
```

## Install Nodemon

```bash
npm install -g nodemon
```

## Running the server

### Go run + Nodemon

```bash
nodemon --exec go run src/main.go --signal SIGTERM
```

Access Swagger API Documentation using this URL:
```url
http://localhost:[your-server-port]/swagger/index.html#/
```

### Docker

- Open docker dekstop, then run this command :

```bash
docker-compose up --build
```

## Tree
This is an example of general folder tree on Fotoyu API service repository.
```bash
.
├── .github
│   └── PULL_REQUEST_TEMPLATE.md
├── docker
│   ├── Dockerfile.dev
│   ├── Dockerfile.loc
│   └── Dockerfile.prd
│   ├── Dockerfile.stg
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
│   │       ├── user
│   │           │── user.go
│   └── main.go
├── .env.example
├── .gitignore
├── README.md
├── docker-compose.yml
└── go.mod
```
