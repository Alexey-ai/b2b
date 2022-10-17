package service

import (
	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/repository"
)

type B2BService struct {
	repo repository.Repository
}

func NewB2BService(repo repository.Repository) *B2BService {
	return &B2BService{repo: repo}
}

func (s *B2BService) GoB2B(request b2b.Request) (int, string, error) {
	return request.Id, request.Value, nil
}
