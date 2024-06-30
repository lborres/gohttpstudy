# build:
# 	@go build -o ./bin/gohttpstudy ./cmd/main.go

dbuild:
	@docker compose up --build -d

build:
	@docker compose up --build