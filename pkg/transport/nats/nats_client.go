package nats

import (
	"github.com/nats-io/nats.go"
)

type Client struct {
	ServerURL string
	*nats.EncodedConn
}

func NewNATSClient(natsServerURL string) *Client {
	return &Client{
		ServerURL: natsServerURL,
	}
}

func (nc *Client) BindReceiverSubject(subject string, subjectChan interface{}) {
	nc.EncodedConn.BindRecvChan(subject, subjectChan)
}

func (nc *Client) BindSenderSubject(subject string, subjectChan chan interface{}) {
	nc.EncodedConn.BindSendChan(subject, subjectChan)
}

// Connect connects to the NATS server
func (nc *Client) Connect() (*Client, error) {
	client, err := nats.Connect(nc.ServerURL)
	if err != nil {
		return nil, err
	}

	if nc.EncodedConn, err = nats.NewEncodedConn(client, nats.JSON_ENCODER); err != nil {
		return nil, err
	}

	return nc, nil
}

// Close closes the connection to the NATS server
func (nc *Client) Close() {
	nc.EncodedConn.Close()

	nc.EncodedConn = nil
}
