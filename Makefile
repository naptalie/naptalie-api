packages:
	go get

tidy:
	go mod tidy

run_server: packages tidy
	go run main.go