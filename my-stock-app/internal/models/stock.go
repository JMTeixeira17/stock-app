package models

import (
	"time"
)

type Stock struct {
	ID         []uint8 `gorm:"primaryKey"`
	Ticker     string  `gorm:"uniqueIndex"`
	TargetFrom string
	TargetTo   string
	Company    string
	Action     string
	Brokerage  string
	RatingFrom string
	RatingTo   string
	Time       time.Time
}
