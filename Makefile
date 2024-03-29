# ==============================================================================
# Dependencies

GOLANG          := golang:1.22
ALPINE          := alpine:3.19

REGISTRY        := registry.digitalocean.com/tacit-tech-registry
APP             := alexandria
BASE_IMAGE_NAME := tacit-tech/alexandria
SERVICE_NAME    := alexandria-api
TAG             := latest
IMAGE_NAME      := $(REGISTRY)/$(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(TAG)

# ==============================================================================
# Containers

docker-build-api:
	docker build --platform=linux/amd64 \
		-f Dockerfile.service \
		-t $(IMAGE_NAME) \
		--build-arg BUILD_REF=$(TAG) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

docker-push-api:
	docker push $(IMAGE_NAME)

# ==============================================================================
# Kubernetes

k8s-apply-api:
	kubectl apply -f k8s

k8s-pods:
	kubectl get pods

k8s-services:
	kubectl get services

k8s-deployments:
	kubectl get deployments

# ==============================================================================
# Development

run:
	go run cmd/alexandria-api/main.go

run-cli:
	go run cmd/alexandria-cli/main.go

ready:
	curl -il http://localhost:8080/readiness

live:
	curl -il http://localhost:8080/liveness
