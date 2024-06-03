BINARY_NAME=app

build:
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}.exe cmd/app/main.go

run:build
	./bin/${BINARY_NAME}.exe

clean:
	go clean
	rm bin/${BINARY_NAME}.exe