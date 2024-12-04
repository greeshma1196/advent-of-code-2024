run:
	go run main.go
build:
	go build -o aoc
test:
	go test ./... -coverprofile cover.out -v
lint: 
	golangci-lint run -v