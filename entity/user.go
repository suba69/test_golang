package entity

import "time"

type user struct {
	id              int
	email           string
	lastEmailSentAt time.Time
}
