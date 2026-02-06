package entity

import "time"

type WeatherQuery struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement"`
	Location            string    `gorm:"size:255;not null;index"`
	Service1Temperature float64   `gorm:"column:service_1_temperature;not null"`
	Service2Temperature float64   `gorm:"column:service_2_temperature;not null"`
	RequestCount        int       `gorm:"column:request_count;not null"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime"`
}
