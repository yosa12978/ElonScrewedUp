package models

import "github.com/google/uuid"

type Tweet struct {
	Id   uuid.UUID
	Url  string
	Date int64
	Text string
}
