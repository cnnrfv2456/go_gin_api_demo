package service

import (
	"errors"
	"go_gin_api_demo/app/front_service/request"
	"go_gin_api_demo/libs/database"
	"go_gin_api_demo/models"
	"go_gin_api_demo/utils"
)

type MemberService struct {
}

func (t *MemberService) CreateMember(request request.CreateMemberRequest) error {
	var (
		member models.Member
	)

	database.Postgres.Where("account = ?", request.Account).First(&member)

	if member.Id != 0 {
		return errors.New("已有重複帳號")
	}

	create := models.Member{
		Name:     request.Name,
		Account:  request.Account,
		Password: utils.StrToMd5(request.Password),
	}

	err := database.Postgres.Create(&create).Error

	if err != nil {
		return errors.New("創建失敗")
	}

	return nil
}
