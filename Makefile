TAG?=1
NAME?=hycdice
REGISTRY=localhost:5001

.PHONY: build
build:
	docker build -t ${REGISTRY}/${NAME}:${TAG} .

.PHONY: push
push:
	docker push ${REGISTRY}/${NAME}:${TAG}
