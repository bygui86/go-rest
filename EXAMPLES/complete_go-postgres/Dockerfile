
FROM golang:1.12-stretch AS gobuilder

WORKDIR /go/src/github.com/bygui86/go-postgres
COPY . .

RUN ["/bin/bash", "-c", "go get -v -d ./..."]
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .

# ---

FROM alpine

RUN apk add --no-cache bash

WORKDIR /bin/
COPY --from=gobuilder /bin/app .

EXPOSE 8080

ENTRYPOINT "/bin/app"
