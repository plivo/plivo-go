.PHONY: build test run

build:
	docker-compose up --build --remove-orphans

test:
	@[ "${CONTAINER}" ] && \
		(docker exec -it $$CONTAINER /bin/bash -c "go test -v -race -cover") || \
		(cd /usr/src/app && go test -v -race -cover)

run:
	@[ "${CONTAINER}" ] && \
		(docker exec -it $$CONTAINER /bin/bash -c 'cd /usr/src/app/go-sdk-test/ && go run test.go') || \
		(cd /usr/src/app/go-sdk-test && go run test.go)