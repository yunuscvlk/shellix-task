.PHONY: server worker

server:
	go run ./cmd/server/server.go

worker:
	go run ./cmd/worker/worker.go