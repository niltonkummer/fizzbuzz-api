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
	ENV=test CONFIG_PATH=./etc/config REDIS_ADDRESS=redis:6379 go test -v `go list ./... | grep -vE 'cmd/|config'` -covermode=count -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html

container-test:
	docker compose -f tests/docker-compose.yml up --build --abort-on-container-exit --remove-orphans --force-recreate

sonar: test
	sonar-scanner \
		-Dsonar.projectKey=FizzBuzz-API \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=sqp_b60e0bc0ba6cbcaa5766fcfd2b1461bfe8d88e23