package server

import (
	"go_gin_api_demo/app/front_service/service"
	"go_gin_api_demo/libs/cache"
	"go_gin_api_demo/libs/database"
	"go_gin_api_demo/repository/postgres"
	"go_gin_api_demo/router"
	"log"

	"github.com/spf13/viper"
)

func init() {
	setYaml()

	db := database.PostgresConn{
		UserName: viper.Get("Postgres.UserName").(string),
		Password: viper.Get("Postgres.Password").(string),
		Host:     viper.Get("Postgres.Host").(string),
		Port:     viper.Get("Postgres.Port").(string),
		Database: viper.Get("Postgres.Database").(string),
	}

	redis := cache.RedisConn{
		Password: viper.Get("Redis.Password").(string),
		Host:     viper.Get("Redis.Host").(string),
		Port:     viper.Get("Redis.Port").(string),
		DB:       viper.Get("Redis.DB").(int),
	}

	//db連線
	db.Connect()

	//redis連線
	redis.Connect()

	var frontService service.Service
	frontService.Repository = postgres.Postgres

	service.FrontService = &frontService
}

//服務啟動
func Run() {
	router.GinRouter()
}

//設置YAML環境變數
func setYaml() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Read config.yaml error: " + err.Error())
	}
}
