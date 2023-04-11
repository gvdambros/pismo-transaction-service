## About

Service responsible for managing accounts and its transactions.

## Technologies

* Go 1.18
* Postgres
* [GORM](https://gorm.io/docs/)
* [ginkgo](https://github.com/onsi/ginkgo)
* [echo](https://github.com/labstack/echo)
* [cpfcnpj](https://github.com/klassmann/cpfcnpj)

## Directory Structure

- `internal`
    - `config`
      - *configuration* yml representation and parsers
    - `domain`
      - application's *business logic*
    - `gen`
      - *generated* files
    - `migration`
      - *database* migration files
    - `pkg`
      - application's *dependencies* (swagger and util modules)
    - `testresources`
      - samples, seeds and fixtures for *testing*

- `it`
    - integration *tests*

- `cmd`
    - *main* files

## Running

### Dependencies
- Local Docker installation

### Makefile targets

| target         | description                     |
|:---------------|:--------------------------------|
| lint           | runs linter (golangci-lint)     |
| tidy           | executes go mod tidy            |
| build-api      | builds api app                  |
| run-api        | runs api app                  |
| test-local     | runs tests                      |
| gen-doc        | codegen for swagger             |
| build-postgres | builds postgres image for tests |
| build-all-deps | builds all test images          |
| clean-all      | cleans all test resources       |

### Debug

To run the service in debugging mode, you will need to:

- Build the postgres image using *make build-postgres*
- Run the postgres image using *make run-postgres*
- Open it in visual studio code and then run using the *.vscode/launch.json* configuration file. 

### Run locally

To run the service in locally, you will need to:

- Build the postgres image using *make build-postgres*
- Run the postgres image using *make run-postgres*
- Build the application using *make build-api*
- Run the application using *make run-api*

### Run tests

To run the tests:

- Run the application using *make test-local*

## API Documentation

To create/update the swagger documentation, you will need to:

- Generate the application documentation files using *make gen-doc*
- Run the application
- Access [Swagger](http://localhost:8080/swagger/index.html#)
