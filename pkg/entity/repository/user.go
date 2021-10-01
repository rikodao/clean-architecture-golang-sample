package repository

import (
	"context"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
)

type User interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
}
