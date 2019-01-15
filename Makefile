PROJECT_REPOSITORY=github.com/becosuke/tasks-api
PROJECT_NAME=tasks
DOCKER_NAME=tasks-golang
DOCKER_GOPATH=/go
PROTO_INCLUDE=-I /usr/include -I protobuf/ -I ${DOCKER_GOPATH}/src -I ${DOCKER_GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
GRPC_PATH=application/grpc
REST_PATH=application/rest
OUTPUT=output
$(eval TIMESTAMP=$(shell date +%Y%m%d%H%M%S))
ifeq ($(shell uname -s),Darwin)
	PRECOMMAND=docker exec -e 'CGO_ENABLED=0' -e 'GOOS=linux' ${DOCKER_NAME}
else
	PRECOMMAND=CGO_ENABLED=0 GOOS=linux
endif

all: dep test build

clean: clean-build clean-test

clean-build:
	${PRECOMMAND} go clean

clean-test:
	${PRECOMMAND} go clean -testcache

dep:
	${PRECOMMAND} dep ensure

dep-update:
	${PRECOMMAND} dep ensure --update

dep-init:
	${PRECOMMAND} dep init

proto: proto-common proto-list proto-context proto-task

proto-common:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/common.proto

proto-list:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/list.proto

proto-context:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/context.proto

proto-task:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/task.proto

test: test-list test-context test-task

test-list:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${REST_PATH}/router/list

test-context:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${REST_PATH}/router/context

test-task:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${REST_PATH}/router/task

build: build-grpc build-rest

build-grpc:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/grpc ${GRPC_PATH}/controller/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=grpc -t tasks-api-grpc:latest -t tasks-api-grpc:${TIMESTAMP}

build-rest:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/rest ${REST_PATH}/controller/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=rest -t tasks-api-rest:latest -t tasks-api-rest:${TIMESTAMP}
