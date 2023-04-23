package controller

import (
	"go_gin_api_demo/app/base"
	"go_gin_api_demo/app/front_service/request"
	"go_gin_api_demo/app/front_service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	base.BaseController
}

// @Tags    會員
// @Summary 創建會員
// @Param   params body request.CreateMemberRequest true "json"
// @accept  application/json
// @Produce application/json
// @produce json
// @Router  /member [post]
// @Success 200 "OK"
func (t *MemberController) CreateMember(ctx *gin.Context) {
	var (
		request request.CreateMemberRequest
		service service.MemberService
	)

	ctx.ShouldBind(&request)

	if errMsg, err := t.ValidateRequest(&request); err != nil {
		t.ResponseJson(ctx, http.StatusBadRequest, errMsg, "")
		return
	}

	err := service.CreateMember(request)

	if err != nil {
		t.ResponseJson(ctx, http.StatusBadRequest, err.Error(), "")
		return
	}

	t.ResponseJson(ctx, http.StatusOK, "", "")
}
