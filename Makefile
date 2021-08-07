APP?=echo
REPO?=docker.pkg.github.com/waynebothaza/echo/
RELEASE?=0.1
VERSION?=0.1.1
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
	@docker build -t ${REPO}${APP}:${VERSION}  .
	@docker tag ${REPO}${APP}:${VERSION} ${REPO}${APP}:${RELEASE}

microk8s: docker
	@docker save ${APP}:${RELEASE} > ${APP}-${RELEASE}.tar
	@microk8s ctr image import ${APP}-${RELEASE}.tar
	@rm ${APP}-${RELEASE}.tar

test:
	@go test -v ./...
