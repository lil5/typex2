default:
	@grep '^[^#[:space:].].*:' Makefile

.PHONEY=test
test:
	go test ./...
coverage:
	 go test -v -coverprofile=profile.cov ./...
build:
	 go build -o typex2
examples:
	make build
	./typex2 -l typescript -i ./examples -o ./examples/typex2.ts
	./typex2 -l dart -i ./examples -o ./examples/typex2.dart
	./typex2 -l swift -i ./examples -o ./examples/typex2.swift
	./typex2 -l kotlin -i ./examples -o ./examples/typex2.kotlin
