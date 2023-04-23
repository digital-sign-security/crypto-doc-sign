.SILENT:

run-be:
	go build backend && go run backend

build-be:
	go build backend

run-fe:
	npm start

build-fe:
	npm build

swagger:
	cd backend && swag init -d "./" -g "cmd/backend/main.go"

run: swagger
	docker-compose up -d --build

migrate-create:
	migrate create -ext sql -dir backend/migrations -seq migration_name
migrate-up:
	docker run -v ./backend/migrations:/backend/migrations --network host migrate/migrate -path=/backend/migrations -database "postgres://postgres:postgres@localhost:5438/postgres?sslmode=disable" up 1
migrate-down:
	docker run -v ./backend/migrations:/backend/migrations --network host migrate/migrate -path=/backend/migrations -database "postgres://postgres:postgres@localhost:5438/postgres?sslmode=disable" down 1
migrate-drop:
