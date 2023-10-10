package data

import (
	externData "github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"os"
)

type Env struct {
	ExternEnv        externData.Env
	PostgresPassword string
	PostgresUser     string
	PostgresHost     string
	PostgresDbname   string
}

func ReadEnv() Env {
	env := Env{
		ExternEnv:        externData.ReadEnv(),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresDbname:   os.Getenv("POSTGRES_DBNAME"),
	}
	return env
}
