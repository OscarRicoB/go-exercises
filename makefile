OUTPUT_NAME=hello-world

build:
	GOARCH=amd64 GOOS=darwin go build -o ${OUTPUT_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${OUTPUT_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${OUTPUT_NAME}-windows main.go

run: build
	./${OUTPUT_NAME}

clean:
	go clean
	rm ${OUTPUT_NAME}-darwin
	rm ${OUTPUT_NAME}-linux
	rm ${OUTPUT_NAME}-windows

test_verbose:
	go test -v ./...

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out