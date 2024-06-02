docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate:
	migrate -path=sql/migrations -database "mysql://application:root@tcp(localhost:3306)/orders" -verbose up

migrate-down:
	migrate -path=sql/migrations -database "mysql://application:root@tcp(localhost:3306)/orders" -verbose down

proto-order:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

grpc:
	evans -r repl

gql-gen:
	go run github.com/99designs/gqlgen generate