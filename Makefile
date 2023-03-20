BIN_NAME=rcon

all: build

build:
	go build -o $(BIN_NAME) ./cmd/cli/rcon-cli.go

clean:
	rm -f $(BIN_NAME)
