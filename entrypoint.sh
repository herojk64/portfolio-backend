#!/bin/sh
set -e
migrate -path /app/sql/migrations -database "$DATABASE_URL" up
exec ./api
