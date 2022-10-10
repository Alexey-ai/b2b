package service

import (
	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/repository"
)

type TestService struct {
	repo repository.Repository
}

func NewTestService(repo repository.Repository) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) GoTest(request b2b.Request) (int, string, error) {
	return request.Id, request.Value, nil
}
