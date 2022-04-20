package service

import (
	"errors"
	"math/rand"

	entity "github.com/epileftro85/goapi/Entity"
	repository "github.com/epileftro85/goapi/Repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repositor repository.PostRepository) PostService {
	repo = repositor
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	var err error = nil

	if post == nil {
		err = errors.New("The post is empty")
	}
	if post.Title == "" {
		err = errors.New("The title is empty")
	}
	if post.Text == "" {
		err = errors.New("The text is empty")
	}

	return err
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
