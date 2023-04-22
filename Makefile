.SILENT:

run-be:
	go build backend && go run backend

build-be:
	go build backend

run-fe:
	npm start

build-fe:
	npm build

swagger-init:
	cd backend && swag init -g /cmd/backend/main.go

complex-run: swagger
	docker-compose up -d --build