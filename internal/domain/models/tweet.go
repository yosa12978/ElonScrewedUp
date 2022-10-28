package models

import "time"

type Tweet struct {
	UUID    string
	Url     string
	Pubdate time.Time
	Text    string
}
