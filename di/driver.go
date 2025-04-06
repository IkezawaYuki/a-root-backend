package di

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
)

var db *gorm.DB
var redisClient *redis.Client

func Connection() {
	db = infrastructure.GetMysqlConnection()
	redisClient = infrastructure.GetRedisConnection()
}

func Close() {
	conn, err := db.DB()
	if err != nil {
		slog.Info("failed to close db")
	} else {
		err := conn.Close()
		if err != nil {
			slog.Info("failed to close db")
		}
	}
	err = redisClient.Close()
	if err != nil {
		slog.Info("failed to close redis")
	}
}

func NewDbDriver() infrastructure.DBDriver {
	return infrastructure.NewDBDriver(db)
}

func NewMailDriver() infrastructure.MailDriver {
	return infrastructure.NewSendgridDriver(
		config.Env.SendgridApiKey,
	)
}
