package dto

import (
	"time"
	"token-assignment/internal/project/entity"
)

type TokenInfo struct {
	Symbol    string     `json:"tokenSymbol"`
	Price     float32    `json:"price"`
	Source    string     `json:"source"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type GetTokenInfoReq struct {
	Symbol    string    `json:"tokenSymbol" validate:"oneof='USTUSD' 'USDC' 'ETH'"`
	Source    string    `json:"source"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type GetTokenInfoResp struct {
	TokenInfo
}

func (resp *GetTokenInfoResp) ToDTO(entResp entity.TokenInfo) {
	resp.Symbol = entResp.Symbol
	resp.Source = entResp.Source
	resp.Price = entResp.Price
	if !entResp.Timestamp.IsZero() {
		resp.Timestamp = &entResp.Timestamp
	}
}
