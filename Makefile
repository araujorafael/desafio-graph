run:
	go run main.go
mocks:
	mockery -all -recursive ./
gqlgen:
	go run github.com/99designs/gqlgen generate