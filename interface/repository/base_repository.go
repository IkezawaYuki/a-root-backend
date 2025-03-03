package repository

import (
	"IkezawaYuki/a-root-backend/infrastructure"
)

type BaseRepository interface {
	Begin() infrastructure.Transaction
}

type baseRepository struct {
	dbDriver infrastructure.DBDriver
}

func NewBaseRepository(dbDriver infrastructure.DBDriver) BaseRepository {
	return &baseRepository{
		dbDriver: dbDriver,
	}
}

func (b *baseRepository) Begin() infrastructure.Transaction {
	return b.dbDriver.Begin()
}
