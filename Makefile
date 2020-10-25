build:
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app cmd/main.go

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

dev:
	@air -c .air.toml