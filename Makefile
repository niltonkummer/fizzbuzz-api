env: # creates and starts the local environment
	docker compose -f docker-compose.yml -p fizz-buzz up -d redis

destroy-env: # destroys the local environment
	docker compose -f docker-compose.yml -p fizz-buzz down --remove-orphans

deps: # installs dependencies
	go install go.uber.org/mock/mockgen@latest

generate:
	go list ./internal/... | xargs -n1 go generate

clean-mocks:
	find ./internal -name "mock_*" -type f -exec rm -rf {} +

mocks: deps clean-mocks generate

test: mocks
	@mkdir -p coverage
	go test -v `go list ./internal/... | grep -vE 'adapters'` -covermode=count -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
