package mock

import (
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	return gormDB, mock, err
}

func NewResult(lastID, rowsAffected int64) driver.Result {
	return sqlmock.NewResult(lastID, rowsAffected)
}
