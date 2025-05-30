#!/bin/sh

dockerize -wait tcp://$DB_HOST:$DB_PORT -timeout 30s

migrate -path $MIGRATION_DIR -database "mysql://$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_SCHEMA?parseTime=true" up

exec "$@"