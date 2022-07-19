run: 
	go build -o bin/main cmd/api/main.go && ./bin/main
run-docker-compose:
	sudo docker-compose up --build
test-all:
	go test ./...
run-mongo-sh: #To connect to mongo from the shell: mongo -u <username> -p <password>
	docker-compose exec mongo sh 