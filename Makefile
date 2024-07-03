export

include .env

LAMBDA_RIE_URL := "https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie"
AWS_LAMBDA_RIE_DIR := ./.aws-lambda-rie
AWS_LAMBDA_RIE_BINARY := $(AWS_LAMBDA_RIE_DIR)/aws-lambda-rie

# Run echo server directly (for local development)
.PHONY: run
run:
	@DISABLE_LAMBDA=1 go run main.go

# Run all test
.PHONY: test
test:
	@go test ./...

.PHONY: setup
setup: install-rie
	go mod tidy

# Install aws-lambda-rie
.PHONY: install-rie
install-rie:
	if [ ! -f "$(AWS_LAMBDA_RIE_BINARY)" ]; then \
    mkdir -p $(AWS_LAMBDA_RIE_DIR); \
    curl -Lo $(AWS_LAMBDA_RIE_BINARY) $(LAMBDA_RIE_URL); \
    chmod +x $(AWS_LAMBDA_RIE_BINARY); \
  fi

