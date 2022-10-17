package service

import (
	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/repository"
)

type B2B interface {
	GoB2B(request b2b.Request) (int, string, error)
}

type Service struct {
	B2B
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		B2B: NewB2BService(*repository.NewRepository()),
	}
}
