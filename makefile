MAIN_FILE = ./cmd/ElonScrewedUp/main.go
BIN_FOLD = ./bin/main.exe

run:
	go mod tidy
	go run $(MAIN_FILE)

build:
	go mod tidy
	go build -o $(BIN_FOLD) $(MAIN_FILE)

deps:
	go mod tidy
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u github.com/google/uuid
	go get -u github.com/gorilla/mux