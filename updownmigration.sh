#!/bin/bash

cd sql/schema || exit 1

# Command 1
echo "Running first command"
goose postgres postgres://fabian:interfeci47@localhost:5432/gator?sslmode=disable down

# Command 2
echo "Running second command"
goose postgres postgres://fabian:interfeci47@localhost:5432/gator?sslmode=disable up

cd ../.. || exit 1