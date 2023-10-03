package storage

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgres() *PostgresStore {
	dsn := "host=localhost user=GO password=PASSWORD dbname=GORM port=5432 sslmode=disable TimeZone=Europe/Zurich"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		sentry.CaptureException(err)
		panic("failed to connect database: " + err.Error())
	}
	err = db.AutoMigrate(&data.Action{})
	if err != nil {
		return nil
	}

	db.Create(&data.Action{
		BucketName:    "production.3s_230913.trimble",
		EquationName:  equation.Addition{}.Name(),
		ListVariables: "KM_000_D,KM_000_G;",
		Active:        true,
	})

	return &PostgresStore{
		db: db,
	}
}
