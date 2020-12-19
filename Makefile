BUILD_DIR = build
CGO_ENABLED ?= 0
GOOS ?= linux

define compile_service
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -ldflags "-s -w" -o ${BUILD_DIR}/$(1) cmd/main.go
endef

define make_docker
	docker build --no-cache --tag=drasko/$(subst docker_,,$(1)) -f docker/Dockerfile .
endef

define make_docker_dev
	docker build --tag=drasko/$(subst docker_dev_,,$(1)) -f docker/Dockerfile.dev ./build
endef

all: gui

.PHONY: gui dockers dockers_dev

clean:
	rm -rf ${BUILD_DIR}

install:
	cp ${BUILD_DIR}/* $(GOBIN)

test:
	GOCACHE=off go test -v -race -tags test $(shell go list ./... | grep -v 'vendor\|cmd')

gui:
	$(call compile_service,$(@))

docker:
	$(call make_docker,$(@))

docker_dev:
	$(call make_docker_dev,$(@))

run:
	cd scripts && ./run.sh
