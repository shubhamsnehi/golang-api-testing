package service

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/shubhamsnehi/golang-api-testing/entity"
	"github.com/shubhamsnehi/golang-api-testing/repository"
)

type PostService interface {
	Validate(post *entity.User) error
	Create1(post *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id string) (*entity.User, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.User) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Name == "" {
		err := errors.New("The post is empty")
		return err
	}

	return nil
}

func (*service) Create1(post *entity.User) (*entity.User, error) {
	uid, _ := uuid.NewRandom()
	post.ID = uid.String()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}

func (*service) FindByID(id string) (*entity.User, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}
