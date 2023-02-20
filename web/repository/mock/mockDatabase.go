package mock

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	mock, sqlMock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mock,
		SkipInitializeWithVersion: true,
	})

	dbMock, err := gorm.Open(dialector)
	if err != nil {
		log.Fatal(err)
	}

	return dbMock, sqlMock
}
