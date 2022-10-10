package service

import (
	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/repository"
)

type Test interface {
	GoTest(request b2b.Request) (int, string, error)
}

type Service struct {
	Test
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Test: NewTestService(*repository.NewRepository()),
	}
}
