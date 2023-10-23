package service

import (
	"pt-test/internal/repo"
)

type Service struct {
	repo *repo.Repository
}

func New(repo *repo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
