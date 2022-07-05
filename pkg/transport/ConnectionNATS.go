package transport

import "gitlab.com/agent-transport-protocol/pkg/transport/nats"

type ConnectionNATS struct {
	Client            nats.Client
	ConnectionChannel chan Connection
	StopChannel       chan bool
}

func NewConnectionNATS() ConnectionNATS {
	connectionNATS := ConnectionNATS{
		Client:            nats.NewNATSClient(),
		ConnectionChannel: make(chan Connection, 1),
		StopChannel:       make(chan bool),
	}

	connectionNATS.Client.BindReceiverSubject("process.agent.connection", connectionNATS.ConnectionChannel)

	connectionNATS.Start()

	return connectionNATS
}

type Connection struct {
}

func (receiver ConnectionNATS) Start() {
connectionReceiver:
	select {
	case connection := <-receiver.ConnectionChannel:
		print(connection)
	case <-receiver.StopChannel:
		break connectionReceiver
	}
}

func (receiver ConnectionNATS) Stop() {
	receiver.StopChannel <- true
}
