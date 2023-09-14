package entity

import "time"

type GetTokenInfoReq struct {
	Symbol    string
	Source    string
	StartTime time.Time
	EndTime   time.Time
}
