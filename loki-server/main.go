package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki"
	"jobscn/ai-flower-pot/loki-server/pkg/vars"
	"jobscn/ai-flower-pot/loki-server/startup"
	"os"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type PostgresConfig struct {
	Host     string
	Port     int32
	User     string
	Password string
	Database string
}

func NewPostgresDBEngine(config *PostgresConfig) *xorm.Engine {
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

	fmt.Println(os.Getenv(`PGSQL_USER`))
	fmt.Println(os.Getenv(`PGSQL_PASSWORD`))
	loki.DBEngine = NewPostgresDBEngine(&PostgresConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     os.Getenv(`PGSQL_USER`),
		Password: os.Getenv("PGSQL_PASSWORD"),
		Database: "loki_server",
	})

	app.Run(vars.ServerAddress)
}
