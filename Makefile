PROJECT_REPOSITORY=github.com/becosuke/tasks-api
PROJECT_NAME=tasks
DOCKER_NAME=tasks-golang
DOCKER_GOPATH=/go
PROTO_INCLUDE=-I protobuf -I /usr/include -I ${DOCKER_GOPATH}/src -I ${DOCKER_GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
GRPC_PATH=application/grpc
REST_PATH=application/rest
TASK_PATH=application/task
OUTPUT=output
TIMESTAMP=$(shell date +%Y%m%d%H%M%S)
PRECOMMAND=docker exec -e 'CGO_ENABLED=0' -e 'GOOS=linux' ${DOCKER_NAME}

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

fmt:
	${PRECOMMAND} sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type f -name *.go -print | xargs -n1 go fmt"

vet:
	${PRECOMMAND} sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type d -print | xargs -IXXX sh -c 'find XXX -maxdepth 1 -type f -name *.go -print | xargs --no-run-if-empty go vet' || :"

lint:
	${PRECOMMAND} sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type f -name *.go -print | xargs -n1 golint -min_confidence=0.8 | grep -v 'should have comment' | grep -v 'ID\\|URL' || :"

proto: proto-common proto-list proto-context proto-task

proto-common:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/common.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/common.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/common.proto

proto-list:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --govalidators_out=${DOCKER_GOPATH}/src protobuf/message/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/list.proto

proto-context:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --govalidators_out=${DOCKER_GOPATH}/src protobuf/message/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/context.proto

proto-task:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --govalidators_out=${DOCKER_GOPATH}/src protobuf/message/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/message/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --go_out=plugins=grpc:${DOCKER_GOPATH}/src protobuf/service/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --grpc-gateway_out=logtostderr=true:${DOCKER_GOPATH}/src protobuf/service/task.proto

swagger: swagger-list swagger-context swagger-task

swagger-list:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/message/list.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/service/list.proto

swagger-context:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/message/context.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/service/context.proto

swagger-task:
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/message/task.proto
	${PRECOMMAND} protoc ${PROTO_INCLUDE} --swagger_out=logtostderr=true:./swagger protobuf/service/task.proto

test: test-grpc

test-grpc: test-list test-context test-task

test-list:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${GRPC_PATH}/controller/list

test-context:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${GRPC_PATH}/controller/context

test-task:
	${PRECOMMAND} go test -v ${PROJECT_REPOSITORY}/${GRPC_PATH}/controller/task

build: build-grpc build-rest

build-grpc:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/grpc ${GRPC_PATH}/main/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=grpc -t tasks-api-grpc:latest -t tasks-api-grpc:${TIMESTAMP}

build-rest:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/rest ${REST_PATH}/main/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=rest -t tasks-api-rest:latest -t tasks-api-rest:${TIMESTAMP}
