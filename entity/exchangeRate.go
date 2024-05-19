package entity

import "time"

type exchangeRate struct {
	iD           int
	exchangeRate float64
	updatedAt    time.Time
}
