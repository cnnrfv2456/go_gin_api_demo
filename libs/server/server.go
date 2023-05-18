package server

import (
	"go_gin_api_demo/libs/cache"
	"go_gin_api_demo/libs/database"
	"go_gin_api_demo/router"
	"log"

	"github.com/spf13/viper"
)

func init() {
	setYaml()
}

//服務啟動
func Run() {
	postgres := database.PostgresConn{
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
	postgres.Connect()
	//redis連線
	redis.Connect()
	//服務啟動
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
