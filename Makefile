gen_docs:
	cd cmd/main && swag init -o "../../docs";

build:
	go build -o bin/bin ./cmd/main/main.go && ./bin/bin;