# Dad Jokes API

This is an example golang api. Super punny.

## Installing and running directly

- Install Go v1.10.x
- Run `go get github.com/golang/dep/cmd/dep`
- Run `dep ensure`
- Run `go run *.go`

## Running with Docker

- `docker-compose up`
- visit `http://localhost:8080`, receive joke

## Building an executable

- Follow the installation steps above, except `go run`
- Run `go build`
- Run `./joke-api`
