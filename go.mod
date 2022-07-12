module github.com/StackVista/agent-transport-protocol

go 1.16

replace github.com/DataDog/sketches-go v1.1.0 => github.com/StackVista/sketches-go v1.1.1

require (
	github.com/DataDog/sketches-go v1.1.0
	github.com/DataDog/zstd v0.0.0-20160706220725-2bf71ec48360
	github.com/nats-io/nats.go v1.16.0
	github.com/stretchr/testify v1.6.1
)

require (
	github.com/gogo/protobuf v1.3.1
	github.com/nats-io/nats-server/v2 v2.8.4 // indirect
)
