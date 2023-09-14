package service

import "token-assignment/internal/project/repository"

type ProjectService interface {
	GetBitfinexTokenInfo(tokenSymbol string) error
}

type ProjectSvc struct {
	repository repository.ProjectRepository
}

var _ ProjectService = (*ProjectSvc)(nil)

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &ProjectSvc{
		repository: repo,
	}
}
