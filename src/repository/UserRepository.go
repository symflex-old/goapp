package repository

import (
	"app/src/entity"
	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	GetUser() *entity.User
}

type UserRepository struct {
	Repository
	Db *pg.DB
}

func (repository *UserRepository) GetUser (id uuid.UUID) *entity.User {
	user := entity.User{Id: id}
	repository.Db.Model(user).First()
	return &user
}