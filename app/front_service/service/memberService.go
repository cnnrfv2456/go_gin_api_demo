package service

import (
	"errors"
	"go_gin_api_demo/app/front_service/request"
	"go_gin_api_demo/utils"
)

func (t *Service) CreateMember(request request.CreateMemberRequest) error {

	members, err := t.Repository.GetMember(map[string]interface{}{"account": request.Account})

	if len(members) > 0 {
		return errors.New("已有重複帳號")
	}

	create := map[string]interface{}{
		"name":     request.Name,
		"account":  request.Account,
		"password": utils.StrToMd5(request.Password),
	}

	err = t.Repository.CreateMember(create)

	if err != nil {
		return errors.New("創建失敗")
	}

	return nil
}
