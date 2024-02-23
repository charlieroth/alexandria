# ==============================================================================
# Define dependencies

GOLANG          := golang:1.22
ALPINE          := alpine:3.19
POSTGRES        := postgres:16.2

NAMESPACE       := alexandria-system
APP             := alexandria
BASE_IMAGE_NAME := tacittech/alexandria
SERVICE_NAME    := alexandria-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)

# ==============================================================================
# Building containers

all: service

docker-build-service:
	docker build \
		-f zarf/docker/Dockerfile.service \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Operations

run:
	go run app/alexandria-server/main.go

admin:
	go run app/tooling/sales-admin/main.go

ready:
	curl -il http://localhost:3000/v1/readiness

live:
	curl -il http://localhost:3000/v1/liveness