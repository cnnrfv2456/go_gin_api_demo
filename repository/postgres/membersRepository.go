package postgres

import "go_gin_api_demo/models"

func (t *Postgres) GetMember(where map[string]interface{}) ([]models.Member, error) {
	var members []models.Member
	err := t.db.Where(where).Find(&members).Error

	return members, err
}

func (t *Postgres) CreateMember(create map[string]interface{}) error {
	var members models.Member

	err := t.db.Model(members).Create(create).Error

	return err
}
