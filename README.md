
# Go REST

Project to explore REST APIs in Golang

---

## TODO list
- [x] simple blog-posts REST APIs for without DB
- [x] testing blog-posts REST APIs
    - [x] unit-testing
    - [x] integration-testing
- [ ] simple products REST APIs for with DB (PostgreSQL)
- [ ] testing products REST APIs
    - [ ] unit-testing
    - [ ] integration-testing

## Build
```shell
make build
```

---

## Test

### Unit-testing
```shell
make unit-test
```

### Integration-testing
```shell
make integr-test
```

### Functional-testing
`TBD`

### Smoke-testing (a.k.a. Build Verification testing)
`TBD`

### Contract-testing
`TBD`

---

## Run

1. start application
```shell
make run
```

2. use [Postman](https://www.postman.com/) to import the [requests collection](./postman) and make some requests 

---

## Prometheus metrics

`TODO`

see https://github.com/prometheus/client_golang/blob/master/prometheus/examples_test.go

---

## Links
### REST
- [x] https://medium.com/@hugo.bjarred/rest-api-with-golang-and-mux-e934f581b8b5
- [x] https://golangcode.com/get-a-url-parameter-from-a-request/
### PostgreSQL
- [ ] https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
### Unit-testing
- [x] https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
- [x] https://golang.org/src/net/http/httptest/example_test.go
- [x] https://github.com/gorilla/mux/issues/373#issuecomment-388568971
### Integration-testing
- [x] https://stackoverflow.com/questions/42474259/golang-how-to-live-test-an-http-server
- [x] https://polothy.github.io/post/2019-04-13-testing-gorrilla-mux-handlers/
- [ ] https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d
- [ ] https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
- [ ] https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html
### Smoke-testing (a.k.a. Build Verification testing)
- [ ] https://medium.com/@felipedutratine/smoke-tests-easily-a008e1e67d50
- [ ] https://www.slideshare.net/alexeygolubev10/smoke-testing-with-go
