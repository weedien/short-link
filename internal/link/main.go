package main

import (
	"context"
	"fmt"
	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/database"
	"shortlink/internal/base/lock"
	"shortlink/internal/base/logging"
	"shortlink/internal/base/mq"
	"shortlink/internal/base/server"
	"shortlink/internal/base/shutdown"
	"shortlink/internal/link/common/config"
	linkservice "shortlink/internal/link/service"
	linktrigger "shortlink/internal/link/trigger/http"
	"syscall"
)

func main() {
	fmt.Println("This is the link-service")

	// 全局日志初始化
	logging.InitLogger()

	// 初始化配置
	config.Init()

	// rocketmq-client 日志
	_ = os.Setenv("rocketmq.client.logRoot", "./logs")
	rmqclient.ResetLogger()

	// 初始化外部依赖
	db := database.ConnectToDatabase()                                             // Postgresql
	rdb := cache.ConnectToRedis()                                                  // Redis
	locker := lock.NewRedisLock(rdb)                                               // DistributedLock - Redis
	eventBus := mq.NewRocketMqBasedEventBus(context.Background(), mq.ProducerMode) // EventBus

	// 创建应用服务
	shortLinkApp := linkservice.NewLinkApplication(db, rdb, locker, eventBus)

	shutdownServer := server.RunHttpServer(func(router fiber.Router) {
		server.NewUriTitleApi(router)
		linktrigger.NewLinkApi(shortLinkApp, router)
		linktrigger.NewLinkRecycleBinApi(shortLinkApp, router)
	})

	shutdown.NewHook().WithSignals(syscall.SIGINT, syscall.SIGTERM).Close(
		// shutdown server
		shutdownServer,
		// shutdown database
		func() {
			if sqlDB, err := db.DB(); err != nil {
				slog.Error("database.DB() failed", "error", err)
			} else {
				if err = sqlDB.Close(); err != nil {
					slog.Error("database.DB().Close() failed", "error", err)
				}
			}
		},
		// shutdown redis
		func() {
			if err := rdb.Close(); err != nil {
				slog.Error("redis.Close() failed", "error", err)
			}
		},
		// shutdown event bus
		func() {
			if eventBus != nil {
				// 事件总线是自己封装的，关闭失败的情况已经在内部进行了处理
				slog.Info("Closing event bus")
				eventBus.Close()
			}
		},
	)
}
