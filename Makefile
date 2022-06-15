test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

.PHONY: test