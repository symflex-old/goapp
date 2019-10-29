package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	User struct {
		tableName struct{} `pg:"user"`
		Id        uuid.UUID `pg:",type:uuid"`
		Email     string
		Password  string
		Name      string
		CreatedAt time.Time
		DeletedAt time.Time
	}

)


