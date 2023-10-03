package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	BucketName    string
	EquationName  string
	ListVariables string
	Active        bool
}
