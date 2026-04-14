package main

import (
	"fmt"

	"github.com/Chihaya-Anon123/TicketHub/internal/config"
	"github.com/Chihaya-Anon123/TicketHub/internal/database"
	"github.com/Chihaya-Anon123/TicketHub/internal/logger"
	"github.com/Chihaya-Anon123/TicketHub/internal/router"
)

func main() {
	//读取配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("load config failed: %v", err))
	}

	//启动 logger
	if err := logger.InitLogger(cfg.Log); err != nil {
		panic(fmt.Sprintf("init logger failed: %v", err))
	}
	defer logger.Sync()

	//初始化并迁移数据库
	if err := database.InitMySQL(cfg.Database); err != nil {
		logger.Log.Fatalf("init mysql failed: %v", err)
	}
	if err := database.AutoMigrate(); err != nil {
		logger.Log.Fatalf("auto migrate failed: %v", err)
	}

	//初始化路由
	r := router.SetupRouter(cfg.JWT)

	//启动服务
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Log.Fatalf("server run failed: %v", err)
	}
	logger.Log.Infow("server starting", "Port", cfg.Server.Port)
}
