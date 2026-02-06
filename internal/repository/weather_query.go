package repository

import (
	"weather-api/internal/entity"

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
	query := entity.WeatherQuery{
		Location:            location,
		Service1Temperature: s1Temp,
		Service2Temperature: s2Temp,
		RequestCount:        count,
	}

	return r.db.Create(&query).Error
}
