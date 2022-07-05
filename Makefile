

proto:
	protoc pkg/model/proto/connection.proto -I vendor -I ${GOPATH}/src -I pkg/model/proto --go_out=${GOPATH}/src