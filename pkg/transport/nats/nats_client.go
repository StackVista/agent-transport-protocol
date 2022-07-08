package nats

import (
	"github.com/nats-io/nats.go"
)

type Client struct {
	ServerURL string
	*nats.EncodedConn
}

func NewNATSClient() (*Client, error) {
	cl := &Client{
		ServerURL: nats.DefaultURL,
	}

	connect, err := cl.Connect()
	if err != nil {
		return nil, err
	}

	cl.EncodedConn = connect

	return cl, nil
}

func (nc *Client) BindReceiverSubject(subject string, subjectChan interface{}) {
	nc.EncodedConn.BindRecvChan(subject, subjectChan)
}

func (nc *Client) BindSenderSubject(subject string, subjectChan chan interface{}) {
	nc.EncodedConn.BindSendChan(subject, subjectChan)
}

// Connect connects to the NATS server
func (nc *Client) Connect() (*nats.EncodedConn, error) {
	client, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	return nats.NewEncodedConn(client, nats.JSON_ENCODER)
}

// Close closes the connection to the NATS server
func (nc *Client) Close() {
	nc.EncodedConn.Close()

	nc.EncodedConn = nil
}
