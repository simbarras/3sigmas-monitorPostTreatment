package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	BucketName    string    `json:"bucketName"`
	EquationName  string    `json:"equationName"`
	ListVariables string    `json:"listVariables"`
	Active        bool      `json:"active"`
}
