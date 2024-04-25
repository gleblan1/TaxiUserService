GOLANGCI_LINT_VERSION := v1.56.2
GOLANGCI_LINT_BIN := ./bin/golangci-lint
GOLANGCI_LINT_CONFIG := ./.golangci.yml

default: lint

$(GOLANGCI_LINT_BIN):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)

lint: $(GOLANGCI_LINT_BIN)
	$(GOLANGCI_LINT_BIN) run --config $(GOLANGCI_LINT_CONFIG)

clean:
	rm -rf ./bin

.DELETE_ON_ERROR:


migrations up:
	$GOPATH/bin/goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
migrations down:
	$GOPATH/bin/goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

build:
	docker build --tag inno-taxi .

