# Fizz-Buzz API Service

[![Build Status](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/build.yml/badge.svg)](https://github.com/niltonkummer/fizzbuzz-api/actions/workflows/build.yml)

## Index
- [Introduction](#introduction)
- [What it is](#what-it-is)
- [How to Use](#how-to-use)
- [Prerequisites](#prerequisites)
- [Running the Service](#running-the-service)
- [API Endpoints](#api-endpoints)
- [Limitations](#limitations)
- [OpenAPI Documentation](#openapi-documentation)
- [How to Test](#how-to-test)
- [Technologies](#technologies)
- [Configuration](#configuration)
- [References](#references)
- [Additional Topics](#additional-topics)
- [Contributing](#contributing)
- [License](#license)

## Introduction
This project implements a scalable and testable Fizz-Buzz microservice in Go. It provides a RESTful API to generate Fizz-Buzz sequences and track usage statistics, designed for extensibility and production-readiness.

## What it is
The Fizz-Buzz service exposes endpoints to:
- Generate a Fizz-Buzz sequence with custom parameters.
- Retrieve statistics about the most frequently requested parameters.

It is built with clean architecture principles, supports in-memory and Redis-based statistics storage, and is ready for containerized deployment.

## How to Use

### Prerequisites
- Go 1.24 or later
- Docker and Docker Compose (for containerized deployment)
- To run bare metal, you need:
  - Go modules enabled
  - Redis (optional, for statistics persistence), you can run it locally or use a managed service
  - Make (for automation)

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

If you want to run with in-memory statistics, set the environment variable `STATS_STORAGE` to `in-memory`. For Redis-based statistics, set it to `redis`.

```sh
ENV=dev STATS_STORAGE=in-memory ./fizzbuzz
```

for Redis, make sure you have a Redis server running and set the environment variable accordingly:

```sh
ENV=dev STATS_STORAGE=redis REDIS_ADDRESS={redis_address} ./fizzbuzz
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
  
For more details on the API, refer to the OpenAPI documentation or look at [http](http) folder

## Limitations
- The param limit for the Fizz-Buzz sequence is set to 500,000 to prevent excessive memory usage.

## Openapi Documentation
The OpenAPI documentation is available at address `localhost:8081`. It provides a detailed description of the API endpoints, request/response formats, and examples.


[Swagger localhost](http://localhost:8081/)
[Swagger online](https://docs.vpneasy.info)


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

### BDD (Behavior-Driven Development)
Feature files and step definitions are in the `tests/` directory.

## Technologies
- Go (Golang)
  - [echo](https://echo.labstack.com/) for the web framework
  - [go-redis](https://github.com/redis/go-redis)
  - [cucumber](https://github.com/cucumber/godog) for BDD
  - [viper](https://github.com/spf13/viper) for configuration management
  - [testify](https://github.com/stretchr/testify) for assertions and mocking
  - [mockegen](https://go.uber.org/mock/mockgen)
- Docker & Docker Compose
- Redis (optional, for stats persistence and for caching)
- Clean Architecture
- Go Modules
- Makefile for automation
- Gherkin for BDD


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
