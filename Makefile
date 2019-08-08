
GO_CMD = go
VARSION = -u

all: tools build

tools:
	$(GO_CMD) get $(VARSION) github.com/google/wire/cmd/wire
	$(GO_CMD) get $(VARSION) github.com/labstack/echo

build:
	$(GO_CMD) build ./cmd/main.go

clean:
	rm -rf main