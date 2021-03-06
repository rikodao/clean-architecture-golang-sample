package repository

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
)

type IUserRepository interface {
	GetUser() (*userModel.UserEntity, error)
	GetUGetUserById(id string) (*userModel.UserEntity, error)
}
