package repository

import (
	"token-assignment/internal/project/entity"

	"github.com/uptrace/bun"
)

type ProjectRepository interface {
	CreateTokenInfo(entity.TokenInfo) error
}

type ProjectRepo struct {
	db *bun.DB
}

var _ ProjectRepository = (*ProjectRepo)(nil)

func NewProjectRepository(db *bun.DB) ProjectRepository {
	return &ProjectRepo{
		db: db,
	}
}
