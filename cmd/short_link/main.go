package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"os/signal"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/config"
	"shortlink/internal/base/database"
	"shortlink/internal/base/lock"
	"shortlink/internal/base/logging"
	"shortlink/internal/base/mq"
	"shortlink/internal/base/server"
	linkservice "shortlink/internal/link/service"
	linktrigger "shortlink/internal/link/trigger/http"
	linkstatsservice "shortlink/internal/link_stats/service"
	linkstatstrigger "shortlink/internal/link_stats/trigger/http"

	"syscall"
)

func main() {
	// 全局日志初始化
	logging.InitLogger()

	// 初始化外部依赖
	db := database.ConnectToDatabase()                            // Postgresql
	rdb := cache.ConnectToRedis()                                 // Redis
	locker := lock.NewRedisLock(rdb)                              // DistributedLock - Redis
	eventBus := mq.NewRocketMqBasedEventBus(context.Background()) // EventBus

	// 创建应用服务
	shortLinkApp := linkservice.NewLinkApplication(db, rdb, locker, eventBus)
	shortLinkStatsApp := linkstatsservice.NewLinkStatsApplication(db)

	shutdownServer := server.RunHttpServerOnPort(config.Port.String(), func(router fiber.Router) {
		server.NewUriTitleApi(router)
		linktrigger.NewLinkApi(shortLinkApp, router)
		linkstatstrigger.NewLinkStatsApi(shortLinkStatsApp, router)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")

	shutdownServer()

	fmt.Println("Running cleanup tasks...")
	// shutdown database
	if sqlDB, err := db.DB(); err != nil {
		slog.Warn("database.DB() failed", "error", err)
	} else {
		err = sqlDB.Close()
		if err != nil {
			slog.Warn("database.DB().Close() failed", "error", err)
		}
	}
	// shutdown redis
	if err := rdb.Close(); err != nil {
		slog.Warn("redis gracefully shutdown failed", "error", err)
	}
	// shutdown event bus
	eventBus.Close() // 事件总线是自己封装的，关闭失败的情况已经在内部进行了处理
	slog.Info("short-link-service was successful shutdown.")
}
