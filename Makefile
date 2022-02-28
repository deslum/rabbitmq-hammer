all: build

go-test:
	go test --race ./... -cover

go-build: go-test
	go build

go-install: go-test go-build
	sudo cp rabbitmq-hammer /usr/local/bin/rabbitmq-hammer

rabbit-up:
	sudo docker-compose down --rmi local
	sudo docker-compose up -d

start:
	go-build
	./rabbitmq-hammer -dep