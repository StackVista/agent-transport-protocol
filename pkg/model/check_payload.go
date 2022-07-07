package model

import "time"

type checkPayload struct {
	messages  []MessageBody
	endpoint  string
	timestamp time.Time
}
