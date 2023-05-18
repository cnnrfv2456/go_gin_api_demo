package postgres

import (
	"go_gin_api_demo/libs/database"

	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres() *Postgres {
	var postgres Postgres
	postgres.db = database.Postgres

	return &postgres
}
