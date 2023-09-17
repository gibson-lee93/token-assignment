package dto

import (
	"math/big"
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
	Symbol    string    `json:"tokenSymbol" validate:"oneof='USDT' 'USDC' 'ETH'"`
	Source    string    `json:"source"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type GetTokenInfoResp struct {
	TokenInfos []TokenInfo
}

func (resp *GetTokenInfoResp) ToDTO(entResp []entity.TokenInfo) {
	resp.TokenInfos = make([]TokenInfo, len(entResp))
	for i, tokenInfo := range entResp {
		resp.TokenInfos[i].Symbol = tokenInfo.Symbol
		resp.TokenInfos[i].Source = tokenInfo.Source
		resp.TokenInfos[i].Price = tokenInfo.Price
		if !tokenInfo.Timestamp.IsZero() {
			resp.TokenInfos[i].Timestamp = &tokenInfo.Timestamp
		}
	}
}

type GetBitfinextTokenResp struct {
	Price     string `json:"last_price"`
	Timestamp string `json:"timestamp"`
}

type ChainLinkTokenResp struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}
