# chimpanzee

This is backend for gorilla (survey app front end)

It uses postgresql as a datastore.

## Assumtions:
You have docker and docker-compose installed on your computer or you have golang compiler and postgresql database.

## Run with docker
`$ docker-compose up`

## Run for local development
1. Download dependencies
`$ go mod download`

2. Export environment variables (they are listed in ".env.example" file)

3. Run the project
`$ go run cmd/api/main.go` 
