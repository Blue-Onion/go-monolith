build: 
	@go build -o bin/api/main cmd/api/main.go
run: build
	@ ./bin/api/main
