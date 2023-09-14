package repository

import "github.com/uptrace/bun"

type ProjectRepository interface{}

type ProjectRepo struct {
	db *bun.DB
}

var _ ProjectRepository = (*ProjectRepo)(nil)

func NewProjectRepository(db *bun.DB) ProjectRepository {
	return &ProjectRepo{
		db: db,
	}
}
