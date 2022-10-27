# TimeMachine API

Written in GoLang, this simple Asynch HTTP-based REST API connects to a PostgreSQL table to return an inventory of products

## Setup

Run a PostgreSQL table
```
docker container run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=timemachinedb postgres:alpine
```

Init the script into PostgreSQL using DataGrip or a PostgreSQL compatible sql client (pgAdmin)

Run the API
```
go run main.go
```

GET the products endpoint
```
url localhost:3100/products
```