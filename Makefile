all: run

get:
	# Get das APIs
	go get github.com/gorilla/mux

run:
	go run *.go
