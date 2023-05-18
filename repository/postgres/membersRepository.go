package postgres

import (
	"go_gin_api_demo/libs/database"
	"go_gin_api_demo/models"
)

func (t *PostgresRepository) GetMember(where map[string]interface{}) ([]models.Member, error) {
	var members []models.Member
	err := database.Postgres.Where(where).Find(&members).Error

	return members, err
}

func (t *PostgresRepository) CreateMember(create map[string]interface{}) error {
	var members models.Member

	err := database.Postgres.Model(members).Create(create).Error

	return err
}
