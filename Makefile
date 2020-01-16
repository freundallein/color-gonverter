export BIN_DIR=bin

export IMAGE_NAME=freundallein/gonverter:latest

init:
	git config core.hooksPath .githooks
run:
	go run main.go
test:
	go test -cover ./service ./converter
build:
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/freundallein/color-gonverter/ protobuf/gonverter.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o $$BIN_DIR/gonverter
dockerbuild:
	docker build -t $$IMAGE_NAME -f Dockerfile .
distribute:
	echo "$$DOCKER_PASSWORD" | docker login -u "$$DOCKER_USERNAME" --password-stdin
	docker build -t $$IMAGE_NAME .
	docker push $$IMAGE_NAME