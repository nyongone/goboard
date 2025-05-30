# Board API with Golang/MySQL

## Start with Docker (Automatically)

```bash
docker-compose build
docker-compose up
```

## Start Manually

```bash
migrate -path MIGRATIONS_DIR -database "mysql://USER:PASS@tcp(HOST:PORT)/SCHEMA?parseTime=true" up

go build -o server main.go
./server
```
