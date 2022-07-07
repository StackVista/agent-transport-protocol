
fmt:
	go fmt ./...

deps:
	go get github.com/DataDog/sketches-go@v1.1.0

proto: deps
	protoc pkg/model/proto/connection.proto -I vendor -I ${GOPATH}/src -I pkg/model/proto --go_out=${GOPATH}/src