package service

import (
	"token-assignment/internal/project/dto"
	"token-assignment/internal/project/repository"
)

type ProjectService interface {
	GetBitfinexTokenInfo(tokenSymbol string) error
	GetTokenInfo(dto.GetTokenInfoReq) (dto.GetTokenInfoResp, error)
	GetChainlinkTokenInfo(tokenSymbol, contractABI, rpcURL string) error
	GetTokenInfoScheduler()
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
