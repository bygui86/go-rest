
# Go REST Postgres

## Build

### From code
```shell
go get -t -d -v ./...
go build -v ./...
```

### From Docker
```shell
docker build . -t go-postgres:latest
```

---

## Run

### Preliminary steps
1. Build application
2. Spin up PostgreSQL with Docker
	```shell
	docker run ...
	```

### From code
```shell
./go-postgres
```

### From Docker\
```shell
docker run -d --name go-postgres -p 8080:8080 go-postgres:latest
```

---

## REST endpoints

* `GET /products` > Fetch a list of products in response to a valid 
* `GET /products/{id}` > Fetch a product in response to a valid 
* `POST /products` > Create a new product in response to a valid 
* `PUT /products/{id}` > Update a product in response to a valid 
* `DELETE /products/{id}` > Delete a product in response to a valid 

---

## Links
* [tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
