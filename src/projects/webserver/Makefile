.PHONY: postgres adminer migrate 

postgres: 
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres:12

adminer: 
	docker run --rm -ti --network host adminer

migrate-up: 
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

migrate-down: 
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down

run-server:
	go run cmd/main.go