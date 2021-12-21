APP?=echo
REPO?=witklippies.azurecr.io
RELEASE?=1.1
VERSION?=1.1.1
USER?=$(shell id -u -n)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_DATE?=$(shell date -u '+%Y-%m-%d')
BUILD_TIME?=$(shell date -u '+%H:%M:%S')

default: ${APP}

clean:
	@rm -f ${APP}

distclean: clean
	@docker image rm ${APP}:${RELEASE} 
	@docker image rm ${APP}:${VERSION} 

${APP}: *.go
	@go build \
		-ldflags "-s -w -X main.BuildVersion=${VERSION} \
		-X main.BuildUser=${USER} \
		-X main.BuildDate=${BUILD_DATE} \
		-X main.BuildTime=${BUILD_TIME} \
		-X main.BuildCommit=${COMMIT}" \
		-o ${APP}

docker:
	@docker build -t ${REPO}/${APP}:${VERSION}  .
	@docker tag ${REPO}/${APP}:${VERSION} ${REPO}/${APP}:${RELEASE}

microk8s: docker
	@docker save ${APP}:${VERSION} > ${APP}-${VERSION}.tar
	@microk8s ctr image import ${APP}-${VERSION}.tar
	@rm ${APP}-${VERSION}.tar

test:
	@go test -v ./...
