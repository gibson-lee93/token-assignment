package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TokenInfo struct {
	bun.BaseModel `bun:"table:tokeninfo"`
	ID            int       `bun:"id,autoincrement,pk"`
	Symbol        string    `bun:"symbol,notnull"`
	Price         float32   `bun:"price,notnull"`
	Source        string    `bun:"srouce,notnull"`
	Timestamp     time.Time `bun:"timestamp,notnull"`
}
