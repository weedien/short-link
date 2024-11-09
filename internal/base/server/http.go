package server

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log/slog"
	"shortlink/internal/base"
	"shortlink/internal/base/server/httperr"
	"strconv"
)

func RunHttpServer(createHandler func(router fiber.Router)) func() {
	port := base.GetConfig().Server.Port
	return RunHttpServerOnPort(strconv.Itoa(port), createHandler)
}

func RunHttpServerOnPort(port string, createHandler func(router fiber.Router)) func() {
	app := fiber.New(fiber.Config{
		AppName:      base.GetConfig().App.Name,
		ErrorHandler: httperr.ErrorHandler,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})
	setupMiddlewares(app)

	// 监控页面
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	basePath := base.GetConfig().Server.BasePath
	if basePath != "" {
		router := app.Group(basePath)
		createHandler(router)
	} else {
		createHandler(app)
	}

	// 处理未找到的路由
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	go func() {
		slog.Info("HTTP server is running on port " + port)
		if err := app.Listen(":" + port); err != nil {
			slog.Error("HTTP server failed to start", "error", err)
		}
	}()

	return func() {
		slog.Info("HTTP server is shutting down")
		if err := app.Shutdown(); err != nil {
			slog.Error("HTTP server failed to shut down", "error", err)
		}
	}
}
