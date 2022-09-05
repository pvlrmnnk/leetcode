.PHONY: new
new:
	@cd ./solution && cookiecutter ../../.cookiecutter/go-solution

.PHONY: lint
lint:
	@golangci-lint run ./...
