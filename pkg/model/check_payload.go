package model

import "time"

type CheckPayload struct {
	Messages  []MessageBody
	Endpoint  string
	Timestamp time.Time
}
