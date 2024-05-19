package entity

import "time"

type exchangeRate struct {
	id           int
	exchangeRate float64
	updatedAt    time.Time
}
