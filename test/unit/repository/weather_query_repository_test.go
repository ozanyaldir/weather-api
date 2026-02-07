package repository_test

import (
	"testing"
	"weather-api/internal/repository"
	"weather-api/test/mock"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestWeatherQueryRepository_Create(t *testing.T) {
	gormDB, mock, err := mock.NewGormMock()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}

	repo := repository.NewWeatherQueryRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `weather_queries`").
		WithArgs("Istanbul", 10.5, 12.5, 1, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create("Istanbul", 10.5, 12.5, 1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
