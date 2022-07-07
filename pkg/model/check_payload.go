package model

import "time"

type CheckPayload struct {
	messages  []MessageBody
	endpoint  string
	timestamp time.Time
}
