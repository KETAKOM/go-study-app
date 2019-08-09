
GO_CMD = go
VARSION = -u
TEST_DIR = ./test/...

all: tools clean fmt tests build

tools:
	$(GO_CMD) get $(VARSION) github.com/google/wire/cmd/wire
	$(GO_CMD) get $(VARSION) github.com/julienschmidt/httprouter

build:
	$(GO_CMD) build -o ./target/hello_world ./cmd/hello_world/main.go
	$(GO_CMD) build -o ./target/todo ./cmd/todo/main.go
	$(GO_CMD) build -o ./target/cmd/api ./cmd/api/main.go

fmt:
	$(GO_CMD) fmt ./...

tests:
	$(GO_CMD) test -v $(TEST_DIR)

clean:
	rm -rf ./target