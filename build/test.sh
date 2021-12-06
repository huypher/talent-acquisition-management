#!/bin/bash

until /usr/bin/pg_isready --quiet --dbname=avay --host=db --port=5432 --username=dbo; do
  sleep 1
done

echo "==> Running unit tests..."
go clean -testcache ./...
go test $(go list ./...) -p 1 --cover
