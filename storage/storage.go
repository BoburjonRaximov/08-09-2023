package storage

import (
	"context"
	"playground/cpp-bootcamp/models"
	"time"
)

type StorageI interface {
	User() UsersI
}
type CacheI interface {
	Cache() RedisI
}

type UsersI interface {
	Create(models.CreateUser) (string, error)
	Update(models.User) (string, error)
	Get(req models.RequestByID) (models.User, error)
	GetByUsername(req models.RequestByUsername) (models.User, error)
	GetAll(models.GetAllUsersRequest) (*models.GetAllUsersResponse, error)
	Delete(req models.RequestByID) (string, error)
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, res interface{}) (bool, error)
	Delete(ctx context.Context, id string) error
}
