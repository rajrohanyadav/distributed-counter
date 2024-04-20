build:
	go build -o distributed-counter

run: build
	./distributed-counter
