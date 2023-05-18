package repository

import "go_gin_api_demo/models"

type Repository interface {
	GetMember(where map[string]interface{}) ([]models.Member, error)
	CreateMember(create map[string]interface{}) error
}
