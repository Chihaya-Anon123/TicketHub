package database

import (
	"fmt"

	"github.com/Chihaya-Anon123/TicketHub/internal/config"
	"github.com/Chihaya-Anon123/TicketHub/internal/logger"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化 MySQL
func InitMySQL(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect database failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get generic database faield: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping database failed: %w", err)
	}

	DB = db
	logger.Log.Infow("mysql connected successfully",
		"host", cfg.Host,
		"port", cfg.Port,
		"dbname", cfg.DBName,
	)
	return nil
}

func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Project{}, &model.ProjectMember{}); err != nil {
		return fmt.Errorf("automigrate database failed: %w", err)
	}

	logger.Log.Infow("database migrate successfully")
	return nil
}
