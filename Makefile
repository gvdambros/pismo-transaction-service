lint:
	golangci-lint run

tidy:
	go mod tidy -compat=1.18

test:
	GO_PROFILE=it \
	ginkgo run -r -p --succinct --output-interceptor-mode=none -coverpkg=./... -coverprofile=coverage.out

test-local: build-all-deps run-all-deps test clean-all

build-api:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o transaction-service ./cmd

run-api:
	GO_PROFILE=it go run ./cmd

build-postgres:
	@./it/config/container/postgres/build_local.sh

build-all-deps: build-postgres

gen-doc:
	@./internal/gen/gen_doc.sh

gen-doc-deps:
	go install github.com/swaggo/swag/cmd/swag@latest

run-postgres:
	docker-compose up postgres -d

run-all-deps:
	docker-compose up -d

clean-postgres:
	docker-compose rm -s -v -f postgres

clean-all:
	docker-compose down
