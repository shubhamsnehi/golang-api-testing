package repository

import "github.com/shubhamsnehi/golang-api-testing/entity"

type PostRepository interface {
	Save(post *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id string) (*entity.User, error)
	Delete(post *entity.User) error
}
