package main

import (
	"os"

	"github.com/Jimeux/fked/internal/api"
	"github.com/Jimeux/fked/internal/domain/fked"
	"github.com/Jimeux/fked/internal/domain/user"
	"github.com/Jimeux/fked/internal/infra/rdbms"
	"github.com/Jimeux/fked/internal/infra/slack"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	MysqlURL   = "FKED_MYSQL_URL"
	SlackToken = "FKED_SLACK_TOKEN"
)

func main() {
	// コンフィグ
	mysqlURL := os.Getenv(MysqlURL)
	slackToken := os.Getenv(SlackToken)

	// DBの初期化
	db, err := xorm.NewEngine("mysql", mysqlURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ワイヤリング
	slackClient := slack.NewClient(slackToken)
	tx := rdbms.NewTx(db)

	userRepository := user.NewRepository()
	fkedRepository := fked.NewRepository()

	fkedService := fked.NewService(tx, fkedRepository, userRepository)

	slackAPI := api.NewSlackAPI(slackClient, fkedService)

	// Ginの初期化
	router := gin.Default()
	router.POST("/slack/event", slackAPI.Event)
	router.Run()
}
