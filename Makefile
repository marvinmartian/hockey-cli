
GOARCH=amd64
GOOS=linux
# GOARM=6
CGO_ENABLED=false
NAME=hockey-cli

.DEFAULT_GOAL := arm
.PHONY: vars test amd64 arm arm6 docker clean clean_build clean_docker

# files: src/main

vars:
	@echo "GOARCH:" ${GOARCH}
	@echo "GOOS:" ${GOOS}
	@echo "GOARM:" ${GOARM}
	@echo "CGO_ENABLED:" ${CGO_ENABLED}
	@echo "NAME:" ${NAME}
	

test:
	go test -cover *.go

arm:
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -a -tags netgo -ldflags '-w' -o ${NAME} main.go

arm6:
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -a -tags netgo -ldflags '-w' -o ${NAME} main.go

container:
	docker build -t ${NAME} .

docker-push: container
	docker push ${NAME}

clean_docker:
	docker rmi ${NAME}

clean_build:
	rm ${NAME}

clean: clean_build clean_docker
	