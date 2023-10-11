package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"idUuid"`
	BucketName    string    `json:"bucketName"`
	EquationName  string    `json:"equationName"`
	ListVariables string    `json:"listVariables"`
	Active        bool      `json:"active"`
}

func ToAction(request AddActionRequest) Action {
	id, err := uuid.Parse(request.ID)
	if err != nil {
		id = uuid.Nil
	}
	return Action{
		ID:            id,
		BucketName:    request.BucketName,
		EquationName:  request.EquationName,
		ListVariables: request.ListVariables,
		Active:        request.Active,
	}
}

type AddActionRequest struct {
	ID            string `json:"idUuid"`
	BucketName    string `json:"bucketName"`
	EquationName  string `json:"equationName"`
	ListVariables string `json:"listVariables"`
	Active        bool   `json:"active"`
}
