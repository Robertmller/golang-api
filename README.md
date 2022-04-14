# golang-api

### Tecnologies used:
- Golang 1.16.7
- Gorilla mux


## How to run:

- `go build cmd/main/main.go`
- `go run cmd/main/main.go`

## Running with Docker:
- `docker build --tag golang-api .`
- `docker run -d -p 8000:8000 golang-api`