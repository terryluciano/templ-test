dev:
	@echo "Starting development server..."
	@air

generate:
	@templ generate

install:
	@go mod tidy
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/air-verse/air@latest
