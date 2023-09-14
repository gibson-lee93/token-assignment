package repository

import (
	"context"
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
