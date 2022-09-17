BIN_NAME=fhirpath

test:
	go test -v ./...

build:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/$(BIN_NAME)

clean:
	rm -rf bin/*
	go clean

