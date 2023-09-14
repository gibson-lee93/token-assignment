package repository

import (
	"context"
	"database/sql"
	"log"
	"token-assignment/internal/project/entity"
)

func (repo *ProjectRepo) CreateTokenInfo(entReq entity.TokenInfo) error {
	if _, err := repo.db.NewInsert().
		Model(&entReq).
		Exec(context.Background()); err != nil {
		log.Println("Failed to create token info")
		log.Println(err)
		return err
	}

	return nil
}

func (repo *ProjectRepo) GetTokenInfo(entReq entity.GetTokenInfoReq) (entResp entity.TokenInfo, err error) {
	query := repo.db.NewSelect().
		Model((*entity.TokenInfo)(nil)).
		Where("symbol = ?", entReq.Symbol)

	// get average price if starttime & enttime is present
	if !entReq.StartTime.IsZero() && !entReq.EndTime.IsZero() {
		avgQuery := repo.db.NewSelect().
			Model((*entity.TokenInfo)(nil)).
			ColumnExpr("AVG(price)").
			Where("symbol = ?", entReq.Symbol).
			Where("timestamp >= ?", entReq.StartTime).
			Where("timestamp <= ?", entReq.EndTime)

		var avgPrice float32
		log.Println(avgQuery)
		if err = avgQuery.Scan(context.Background(), &avgPrice); err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return entResp, err
		}

		entResp.Symbol = entReq.Symbol
		entResp.Price = avgPrice
		entResp.Source = entReq.Source
		return entResp, nil
	}

	// return latest token if starttime & enttime is not present
	if err = query.
		Order("timestamp DESC").
		Limit(1).
		Scan(context.Background(), &entResp); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return entResp, err
	}

	return entResp, nil
}
