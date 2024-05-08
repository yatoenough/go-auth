BINARY_NAME=app

build:
	go build -o bin/${BINARY_NAME}.exe cmd/app/main.go

run:build
	./bin/${BINARY_NAME}.exe

clean:
	go clean
	rm ${BINARY_NAME}.exe
