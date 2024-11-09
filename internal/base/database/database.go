package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"shortlink/internal/base"
)

func ConnectToDatabase() *gorm.DB {

	config := base.GetConfig().Database

	dsn := config.Dsn
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	// Setup sharding
	if config.EnableSharding {
		setupSharding(db)
	}
	return db
}
