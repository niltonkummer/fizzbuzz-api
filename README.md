# Fizz-Buzz API Service

[![Test Status](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/test.yml/badge.svg)](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/test.yml)
[![Build Status](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/build.yml/badge.svg)](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/build.yml)

## Introduction
This project implements a scalable and testable Fizz-Buzz microservice in Go. It provides a RESTful API to generate Fizz-Buzz sequences and track usage statistics, designed for extensibility and production-readiness.

## What it is
The Fizz-Buzz service exposes endpoints to:
- Generate a Fizz-Buzz sequence with custom parameters.
- Retrieve statistics about the most frequently requested parameters.

It is built with clean architecture principles, supports in-memory and Redis-based statistics storage, and is ready for containerized deployment.

## How to Use

### Running the Service
You can run the service locally using Docker Compose:

```sh
docker-compose up --build
```

Or build and run directly with Go:

```sh
go build -o fizzbuzz ./cmd/api/main.go
```

Then execute the binary:

```sh
ENV=dev ./fizzbuzz
```

The service will start on the port defined in `etc/config/server.dev.env` (default: 8080).

### API Endpoints

#### Generate Fizz-Buzz
- **POST** `/fizzbuzz`
- **Body Example:**
  ```json
  {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "Fizz",
    "str2": "Buzz"
  }
  ```
- **Response Example:**
  ```json
  {
    "result": "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
  }
  ```

#### Get Statistics
- **GET** `/stats`
- **Response Example:**
  ```json
  {
    "parameters": {
      "int1": 3,
      "int2": 5,
      "limit": 15,
      "str1": "Fizz",
      "str2": "Buzz"
    },
    "hits": 42
  }
  ```

## How to Test

### Unit Tests
Run all unit tests with:

```sh
make test
```

Or directly with Go:

```sh
go test ./...
```

### Test Coverage
Generate a coverage report:

```sh
make coverage
```

View the HTML report:

```sh
open coverage/coverage.html
```

### BDD (Behavior-Driven Development)
Feature files and step definitions are in the `tests/` directory. Run BDD tests with:

```sh
make bdd
```

## Technologies
- Go (Golang)
- Docker & Docker Compose
- Redis (optional, for stats persistence and for caching)
- Clean Architecture
- Go Modules
- Makefile for automation
- Ginkgo/Gomega (for BDD)

## Configuration
- Environment variables are managed in `etc/config/server.${ENV}.env`.
- You can switch stats storage between in-memory and Redis in the configuration.

## References
- [Go Documentation](https://golang.org/doc/)
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Docker Documentation](https://docs.docker.com/)

## Additional Topics

### Project Structure
- `cmd/api/`: Main entrypoint for the API server
- `internal/`: Application logic, adapters, and domain models
- `config/`: Configuration loading
- `tests/`: BDD and integration tests
- `coverage/`: Test coverage reports

### Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

### License
This project is licensed under the MIT License.
