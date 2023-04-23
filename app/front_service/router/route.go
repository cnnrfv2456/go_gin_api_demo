package router

import (
	"go_gin_api_demo/app/front_service/controller"

	"github.com/gin-gonic/gin"
)

// Api
// @title                      gin_deom
// @version                    1.0
// @description                前台api
// @BasePath                   /api/front
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func Route(r *gin.Engine) {
	api := r.Group("api/front")

	api.POST("/member", new(controller.MemberController).CreateMember) //創建會員

}
