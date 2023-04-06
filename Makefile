build:
	go build cmd/main.go
	./main.exe

run:
	go run cmd/main.go

test:
	go test -v

up:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' up

down:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' down