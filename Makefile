project_name = app
image_name = app:latest

include .env
export

POSTGRES_IMAGE_NAME = postgres:latest
POSTGRES_CONTAINER_NAME = postgres-db
run-local:
	go fmt ./... && gosec ./... && air main.go

docs-generate:
	swag init

gen-erc20-binding:
	abigen --abi=./internal/smartcontract/ABI/ERC20.abi.json --pkg=ERC20 --out=./internal/smartcontract/binding/ERC20/ERC20.go

gen-pool-swap-binding:
	abigen --abi=./internal/smartcontract/ABI/PoolSwap.abi.json --pkg=PoolSwap --out=./internal/smartcontract/binding/PoolSwap/PoolSwap.go

gen-bridge-binding:
	abigen --abi=./internal/smartcontract/ABI/Bridge.abi.json --pkg=Bridge --out=./internal/smartcontract/binding/Bridge/Bridge.go

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

build:
	docker build -t $(image_name) .

stop-postgres:
	docker stop $(POSTGRES_CONTAINER_NAME)
	docker rm $(POSTGRES_CONTAINER_NAME)

# Start the PostgreSQL container
start-postgres:
	docker run --name $(POSTGRES_CONTAINER_NAME) --env-file .env -d -p 15432:5432 $(POSTGRES_IMAGE_NAME)

start:
	make start-postgres
	docker run -d -p 8080:8080 --env-file .env --link $(POSTGRES_CONTAINER_NAME):postgres $(image_name)

build-no-cache:
	docker build --no-cache -t $(image_name) .

service-stop:
	docker compose down

service-start:
	make stop
	docker compose up
