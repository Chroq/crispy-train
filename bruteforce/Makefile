# Makefile for building all Go commands in cmd/*

# Variables
GO_CMD := go
CMD_DIR := cmd
BIN_DIR := bin

# Targets
.PHONY: all clean build

build:
	@echo "Building all sources in cmd..."
	@mkdir -p $(BIN_DIR)
	@touch $(BIN_DIR)/.gitignore
	@go build -o $(BIN_DIR)/combination_async cmd/combination_async/main.go
	@go build -o $(BIN_DIR)/async cmd/async/main.go

benchmarks:
	@echo "Running benchmarks..."
	@hyperfine -N ./$(BIN_DIR)/async ./$(BIN_DIR)/combination_async

all: build benchmarks

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)