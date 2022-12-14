# TimeMachine API

Written in GoLang, this simple Asynch HTTP-based REST API connects to a PostgreSQL table to return an inventory of products

## Setup

Run a PostgreSQL table using podman

What is podman? Quickly, podman is an open source direct replacement of docker, however it is daemon-less and root-less
```
❯ podman container run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=timemachinedb postgres:alpine
❯ podman container ls
CONTAINER ID  IMAGE                              COMMAND     CREATED        STATUS            PORTS                   NAMES
177b5797480e  docker.io/library/postgres:alpine  postgres    2 minutes ago  Up 2 minutes ago  0.0.0.0:5432->5432/tcp  postgres
```

Init the script into PostgreSQL using DataGrip or a PostgreSQL compatible sql client (pgAdmin)

## Running the API

Run the API
```
go run main.go
```

GET the products endpoint
```
❯ curl localhost:3100/products                                                                                                                                                            ~
[{"id":1,"productName":"ThinkPad T Series","productType":"laptop","supplierCost":"800.00","salePrice":"1000.99"},{"id":2,"productName":"Seiko 5","productType":"watch","supplierCost":"45.00","salePrice":"99.95"},{"id":3,"productName":"Das Beer","productType":"beer","supplierCost":"3","salePrice":"6.00"}]
```
