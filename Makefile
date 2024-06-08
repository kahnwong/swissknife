EXECUTABLE_NAME := swissknife
BUILD_DIR := build

.PHONY: all clean

all: clean windows-amd64 darwin-amd64 darwin-arm64 linux-amd64 linux-arm64

clean:
	rm -rf $(BUILD_DIR)

windows-amd64: $(SRC)
	GOOS=windows GOARCH=amd64 go build -o build/$(EXECUTABLE_NAME)-windows-amd64.exe $(SRC)

# darwin-amd64: $(SRC)
# 	GOOS=darwin GOARCH=amd64 go build -o build/$(EXECUTABLE_NAME)-darwin-amd64 $(SRC)

darwin-arm64: $(SRC)
	GOOS=darwin GOARCH=arm64 go build -o build/$(EXECUTABLE_NAME)-darwin-arm64 $(SRC)

linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o build/$(EXECUTABLE_NAME)-linux-amd64 $(SRC)

linux-arm64: $(SRC)
	GOOS=linux GOARCH=arm64 go build -o build/$(EXECUTABLE_NAME)-linux-arm64 $(SRC)
# --------
test:
	go test ./...
