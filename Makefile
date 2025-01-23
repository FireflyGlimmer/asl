ASL_PROJECT := $(shell pwd)
OUT_DIR := $(ASL_PROJECT)/out
PROTO_DIR := $(ASL_PROJECT)/proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

clean:
	@mkdir -p $(OUT_DIR)
	@rm -rf $(OUT_DIR)/*

all: linux-release windows-release darwin-release

windows-release: clean
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o $(OUT_DIR)/aslgo_windows_arm64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/aslgo_windows_amd64.exe

linux-release: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(OUT_DIR)/aslgo_linux_arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/aslgo_linux_amd64

darwin-release: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/aslgo_darwin_arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/aslgo_darwin_amd64
