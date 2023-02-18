package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable struct {
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     string
	MysqlDatabase string
	Port          string
	AuthorizerUrl string
}

var Env *EnvironmentVariable

func GetEnvironmentVariables() *EnvironmentVariable {
	godotenv.Load(".env")

	Env = &EnvironmentVariable{
		MysqlUser:     os.Getenv("MYSQL_USER"),
		MysqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MysqlHost:     os.Getenv("MYSQL_HOST"),
		MysqlPort:     os.Getenv("MYSQL_PORT"),
		MysqlDatabase: os.Getenv("MYSQL_DATABASE"),
		Port:          os.Getenv("PORT"),
		AuthorizerUrl: os.Getenv("AUTHORIZER_URL"),
	}
	return Env
}
