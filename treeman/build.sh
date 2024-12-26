#!/bin/bash

# Build the frontend
cd frontend
pnpm run build
cd ..

# Build the server
go build -o ./tmp/main ./main.go
