package model

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeZeroTimestamp(t *testing.T) {
	header := MessageHeader{
		Version:        MessageV3,
		Encoding:       MessageEncodingZstdPB,
		Type:           TypeCollectorProc,
		SubscriptionID: 0,
		OrgID:          0,
		Timestamp:      0,
	}
	headerBytes, err := encodeHeader(header)
	assert.NoError(t, err)
	headerB64 := base64.StdEncoding.EncodeToString(headerBytes)

	// the same values are expected in the StackState receiver
	// make sure of backward compatibility when changing it
	assert.EqualValues(t, "AwIMAAAAAAAAAAAAAAAAAA==", headerB64)
}

func TestEncodeNonZeroTimestamp(t *testing.T) {
	header := MessageHeader{
		Version:        MessageV3,
		Encoding:       MessageEncodingZstdPB,
		Type:           TypeCollectorProc,
		SubscriptionID: 0,
		OrgID:          0,
		Timestamp:      1638527655412,
	}
	headerBytes, err := encodeHeader(header)
	assert.NoError(t, err)
	headerB64 := base64.StdEncoding.EncodeToString(headerBytes)

	// the same values are expected in the StackState receiver
	// make sure of backward compatibility when changing it
	assert.EqualValues(t, "AwIMAAAAAAAAAAF9f9vd9A==", headerB64)
}

func TestEncodeDecodeMessage(t *testing.T) {
	header := MessageHeader{
		Version:        MessageV3,
		Encoding:       MessageEncodingProtobuf,
		Type:           TypeCollectorConnections,
		SubscriptionID: 0,
		OrgID:          0,
		Timestamp:      1638527655412,
	}
	body := CollectorConnections{
		HostName: "hostname",
		Connections: []*Connection{
			{
				Pid: 0,
				Laddr: &Addr{
					Host: &Host{
						Id:          0,
						OrgId:       0,
						Name:        "",
						Tags:        nil,
						AllTags:     nil,
						NumCpus:     0,
						TotalMemory: 0,
					},
					Ip:   "10.0.0.0",
					Port: 80,
				},
				Raddr:                  nil,
				BytesSentPerSecond:     0,
				BytesReceivedPerSecond: 0,
				Family:                 0,
				Type:                   0,
				PidCreateTime:          0,
				Namespace:              "",
				Direction:              0,
				ConnectionIdentifier:   "",
				ApplicationProtocol:    "",
				Metrics:                nil,
			},
		},
	}
	encodeMessage, err := EncodeMessage(Message{
		Header: header,
		Body:   &body,
	},
	)
	assert.NoError(t, err)
	fmt.Println(encodeMessage)

	decodedMessage, err := DecodeMessage(encodeMessage)
	assert.NoError(t, err)
	connections := decodedMessage.Body.(*CollectorConnections)
	assert.EqualValues(t, "10.0.0.0", connections.Connections[0].Laddr.Ip)

}
