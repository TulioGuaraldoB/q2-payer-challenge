package db

import (
	"fmt"
	"log"

	"github.com/TulioGuaraldoB/q2-payer-challenge/config/env"
	"github.com/TulioGuaraldoB/q2-payer-challenge/infra/db/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func parseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.Env.MysqlUser,
		env.Env.MysqlPassword,
		env.Env.MysqlHost,
		env.Env.MysqlPort,
		env.Env.MysqlDatabase,
	)
}

func StartDatabase() {
	dsn := parseDSN()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect to MySql. %s", err.Error())
		log.Fatal(errMessage)
	}

	db = database

	migration.Run(db)
}

func OpenConnection() *gorm.DB {
	if db == nil {
		StartDatabase()
	}

	return db
}
