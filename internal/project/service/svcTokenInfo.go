package service

import (
	"token-assignment/internal/project/dto"
	"token-assignment/internal/project/entity"
)

func (svc *ProjectSvc) GetTokenInfo(dtoReq dto.GetTokenInfoReq) (dtoResp dto.GetTokenInfoResp, err error) {
	entReq := entity.GetTokenInfoReq{
		Symbol:    dtoReq.Symbol,
		Source:    dtoReq.Source,
		StartTime: dtoReq.StartTime,
		EndTime:   dtoReq.EndTime,
	}

	entResp, err := svc.repository.GetTokenInfo(entReq)
	if err != nil {
		return dtoResp, err
	}

	dtoResp.ToDTO(entResp)
	return dtoResp, nil
}
