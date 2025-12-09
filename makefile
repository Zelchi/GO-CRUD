PROJECT_NAME=API_GO

run: 
	go run ./cmd/main.go

build: binaries
	GOOS=linux GOARCH=amd64 go build -o ./binaries/${PROJECT_NAME} ./cmd/main.go

clean:
	go clean
	rm -rf ./binaries/**

binaries:
	mkdir -p binaries