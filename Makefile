BINARY=teamguru

build:
	go build -o ${BINARY}

install:
	go install

test:
	go test ./...

clean:
	go clean
