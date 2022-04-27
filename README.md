# golang-api

### Tecnologies used:
- Golang 1.16.7
- Gorilla mux
- GORM (ORM)
- GRPC
- KAFKA


## How to run:

- `go build cmd/main/main.go`
- `go run cmd/main/main.go`

## Running with Docker:
- `docker build --tag golang-api .`
- `docker run -d -p 8000:8000 golang-api`

## Routes:
- `/movies/` - GET ALL MOVIES
- `/movie/id` - GET ONE MOVIE
- `/movie/` -  POST ONE MOVIE
- `/movie/id` - PUT ONE MOVIE
- `/movie/id` - DELETE ONE MOVIE