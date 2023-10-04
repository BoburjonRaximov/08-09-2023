package storage

import "playground/cpp-bootcamp/models"

type StorageI interface {
	User() UsersI
}
type UsersI interface {
	Create(models.CreateUser) (string, error)
	Update(models.User) (string, error)
	Get(req models.RequestByID) (models.User, error)
	GetByUsername(req models.RequestByUsername) (models.User, error)
	GetAll(models.GetAllUsersRequest) (*models.GetAllUsersResponse, error)
	Delete(req models.RequestByID) (string, error)
	ChangePassword(req models.ChangePassword) (string, error)
}
