build:
	go build -o bin/bin cmd/main/main.go && ./bin/bin ./config/local.config.yaml