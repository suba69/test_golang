package entity

import "time"

type subscribedUser struct {
	id              int
	email           string
	lastEmailSentAt time.Time
}
