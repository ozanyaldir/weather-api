package models

import "time"

type WeatherQuery struct {
	ID                  uint      `gorm:"primaryKey"`
	Location            string    `gorm:"size:255;not null;index"`
	Service1Temperature float64   `gorm:"not null"`
	Service2Temperature float64   `gorm:"not null"`
	RequestCount        int       `gorm:"not null"`
	CreatedAt           time.Time `gorm:"index"`
}
