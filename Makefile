BINARY_NAME=app

build:
	go build -o bin/${BINARY_NAME}.exe cmd/app/main.go

run:
	build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}.exe
