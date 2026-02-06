package repository

import (
	"time"
	"weather-api/internal/model"

	"gorm.io/gorm"
)

type IWeatherQueryRepository interface {
	Create(location string, s1Temp, s2Temp float64, count int) error
}

type WeatherQueryRepository struct {
	db *gorm.DB
}

func NewWeatherQueryRepository(db *gorm.DB) *WeatherQueryRepository {
	return &WeatherQueryRepository{
		db: db,
	}
}

func (r *WeatherQueryRepository) Create(location string, s1Temp, s2Temp float64, count int) error {
	query := model.WeatherQuery{
		Location:            location,
		Service1Temperature: s1Temp,
		Service2Temperature: s2Temp,
		RequestCount:        count,
		CreatedAt:           time.Now(),
	}

	return r.db.Create(&query).Error
}
