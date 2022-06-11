# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
MAIN="cmd/15to18.go"
BINARY_NAME="bin/15to18"
FILE_NAME=15to18

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)

tests:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -v -cover ./...

run:
	$(GORUN) $(MAIN) $(ARGS)

compile:
	# 64-Bit
	# MacOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ${BINARY_NAME}-darwin-amd64 $(MAIN)
	# MacOS
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ${BINARY_NAME}-darwin-arm64 $(MAIN)
	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ${BINARY_NAME}-linux-amd64 $(MAIN)
	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ${BINARY_NAME}-windows-amd64 $(MAIN)

clean:
	rm bin/${FILE_NAME}-darwin-amd64
	rm bin/${FILE_NAME}-darwin-arm64
	rm bin/${FILE_NAME}-linux-amd64
	rm bin/${FILE_NAME}-windows-amd64