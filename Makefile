BIN_NAME=rcon-cli

all: build

build:
	go build -o build/$(BIN_NAME) ./cmd/cli/rcon-cli.go

multi-arch:
	scripts/build-multi-arch.sh build/$(BIN_NAME)

clean:
	rm -f $(BIN_NAME)*
