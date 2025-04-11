include secret.env

up:
	docker-compose up -d --build

down:
	docker-compose down

stop:
	docker-compose stop

migrateup:
	migrate -path db/migrations -database "postgres://smarchz:${DB_PASSWORD}@localhost:5432/secretaria?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://smarchz:${DB_PASSWORD}@localhost:5432/secretaria?sslmode=disable" -verbose down

backup:
	docker exec -t postgres14 pg_dumpall -c -U smarchz > db/backup/backup_$(date +%Y-%m-%d"_"%H.%M.%S).sql

protoc:
	protoc internal/adapters/driving/grpc/proto/finance.proto --go_out=internal/adapters/driving/grpc --go-grpc_out=internal/adapters/driving/grpc

.PHONY: up down stop migrateup migratedown backup protoc