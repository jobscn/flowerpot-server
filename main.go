package main

import (
	"fmt"
	"jobscn/ai-flower-pot/startup"
	"jobscn/ai-flower-pot/util"
	"jobscn/ai-flower-pot/ymer"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type PostgresConfig struct{
	Host string
	Port int32
	User string
	Password string
	Database string
}

func NewPostgresDBEngine(config *PostgresConfig) *xorm.Engine{
	// datasource arg format
	dataSourceArg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	// init postgres database engine
	engine, err := xorm.NewEngine("postgres", dataSourceArg)
	if err != nil {
		panic(err)
	}

	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.ShowSQL(true)

	err = engine.Ping()
	if err != nil {
		panic(err)
	}

	return engine
}

func main() {
	app := gin.Default()

	startup.RegisterGinRouter(app)

	ymer.DBEngine = NewPostgresDBEngine(&PostgresConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		Database: "loki_server",
	})

	app.Run(util.ServerAddress)
}
