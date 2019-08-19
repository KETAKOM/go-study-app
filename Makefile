
GO_CMD = go
VARSION = -u
TEST_DIR = ./test/...

all: tools clean fmt tests build

tools:
	$(GO_CMD) get $(VARSION) github.com/google/wire/cmd/wire
	$(GO_CMD) get $(VARSION) github.com/julienschmidt/httprouter
	$(GO_CMD) get $(VARSION) github.com/go-sql-driver/mysql
	$(GO_CMD) get $(VARSION) github.com/aws/aws-sdk-go
	$(GO_CMD) get $(VARSION) github.com/aws/aws-sdk-go/service/sqs

build:
	$(GO_CMD) build -o ./target/hello_world ./cmd/hello_world/main.go
	$(GO_CMD) build -o ./target/todo ./cmd/todo/main.go
	$(GO_CMD) build -o ./target/cmd/api ./cmd/api/main.go
	$(GO_CMD) build -o ./target/test_wire ./cmd/test_wire/main.go

fmt:
	$(GO_CMD) fmt ./...

wire:
	rm -rf ./registory/wire_gen.go
	wire gen ./registory/...

tests:
	$(GO_CMD) test -v $(TEST_DIR)

mock-db: #Starting the docker daemon
	docker container run --rm -d \
		-v $(PWD)/database/init:/docker-entrypoint-initdb.d \
		-e MYSQL_ROOT_PASSWORD=root \
		-p 43306:3306 --name mysql mysql:5.7

mock-db-stop:
	docker container stop mysql

clean:
	rm -rf ./target