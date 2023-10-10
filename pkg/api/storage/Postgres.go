package storage

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgres(env data.Env) *PostgresStore {
	//dsn := "host=localhost user=GO password=PASSWORD dbname=GORM port=5432 sslmode=disable TimeZone=Europe/Zurich"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Zurich", env.PostgresHost, env.PostgresUser, env.PostgresPassword, env.PostgresDbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		sentry.CaptureException(err)
		panic("failed to connect database: " + err.Error())
	}
	err = db.AutoMigrate(&data.Action{})
	if err != nil {
		sentry.CaptureException(err)
		panic("failed to migrate database: " + err.Error())
	}

	return &PostgresStore{
		db: db,
	}
}

func (p *PostgresStore) GetAllActions() []data.Action {
	var actions []data.Action
	p.db.Find(&actions)
	return actions
}

func (p *PostgresStore) AddAction(action data.Action) {
	p.db.Create(&action)
}

func (p *PostgresStore) UpdateAction(action data.Action) {
	p.db.Save(&action)
}

func (p *PostgresStore) DeleteAction(action data.Action) {
	p.db.Delete(&action)
}

func (p *PostgresStore) FindAction(id string) data.Action {
	var action data.Action
	uuidId := uuid.MustParse(id)
	p.db.First(&action, uuidId)
	return action
}

func (p *PostgresStore) FindActionsByBucket(name string) []data.Action {
	var actions []data.Action
	p.db.Where("active = true AND bucket_name = ?", name).Find(&actions)
	return actions
}
