
GO_CMD = go
VARSION = -u

all: tools clean fmt build

tools:
	$(GO_CMD) get $(VARSION) github.com/google/wire/cmd/wire
	$(GO_CMD) get $(VARSION) github.com/labstack/echo

build:
	$(GO_CMD) build -o ./target/hello_world ./cmd/hello_world/main.go
	$(GO_CMD) build -o ./target/todo ./cmd/todo/main.go

fmt:
	$(GO_CMD) fmt ./...

clean:
	rm -rf ./target