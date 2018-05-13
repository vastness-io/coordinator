VERSION=$(shell cat ./VERSION)
COMMIT=$(shell git rev-parse --short HEAD)
LATEST_TAG=$(shell git tag -l | head -n 1)

export VERSION COMMIT LATEST_TAG
.PHONY: test

test:
	@echo "=> Running tests"
	./hack/run-tests.sh

build:
	rm -rf ./bin
	./hack/cross-platform-build.sh

verify:
	./hack/verify-version.sh

up: build
	docker build -t vastness.io/coordinator:${VERSION} .
	docker-compose up

generate:
	@echo "=> generating mocks"
	./hack/generate-mocks.sh

container: build
	docker build -t quay.io/vastness/coordinator:${COMMIT} .

push: container
	docker push quay.io/vastness/coordinator:${COMMIT}
	docker tag quay.io/vastness/coordinator:${COMMIT} quay.io/vastness/coordinator:${VERSION}
	docker push quay.io/vastness/coordinator:${VERSION}
	docker tag quay.io/vastness/coordinator:${COMMIT} quay.io/vastness/coordinator:latest
	docker push quay.io/vastness/coordinator:latest
