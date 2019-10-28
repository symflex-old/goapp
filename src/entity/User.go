package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id uuid.UUID `pg:",pk,type:uuid default uuid_generate_v4()"`
	Email string
	Password string
	Name string
	CreatedAt time.Time `pg: "default:now()"`
	DeletedAt time.Time
}
