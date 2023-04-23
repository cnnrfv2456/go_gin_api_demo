package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConn struct {
	UserName, Password, Host, Port, Database string
}

var Postgres *gorm.DB

func (t *PostgresConn) Connect() {

	//組合sql連線字串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", t.Host, t.UserName, t.Password, t.Database, t.Port)

	var err error

	//連接MySQL
	Postgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})

	if err != nil {
		log.Panic("Connection to postgres failed: " + err.Error())
	}

	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	conn, err := Postgres.DB()

	if err != nil {
		log.Panic("Get db failed: " + err.Error())
	}
	conn.SetConnMaxLifetime(24 * time.Hour)
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(20)
}
