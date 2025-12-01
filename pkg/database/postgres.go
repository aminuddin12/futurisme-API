package database

import (
	"fmt"
	"log"
	"time"

	"futurisme-api/config"
	"futurisme-api/internal/modules/user/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDatabase(cfg *config.Config) *gorm.DB {
	// FIX: Menambahkan cfg.Database.Port yang sebelumnya terlewat
	// Urutan harus sesuai dengan placeholder %s: host, user, password, dbname, port, sslmode, timezone
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Name,
		cfg.Database.Port, // <-- Sebelumnya baris ini hilang
		cfg.Database.SSLMode,
		cfg.Database.TimeZone,
	)

	var gormLogger logger.Interface
	if cfg.App.Debug {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// HYBRID MIGRATION LOGIC
	if cfg.App.Env == "dev" {
		log.Println("⚠️  Running in DEV mode: GORM AutoMigrate is ENABLED")

		err := db.AutoMigrate(&entity.User{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}

		log.Println("✅ Database migrated successfully")
	} else {
		log.Println("🔒 Running in PROD mode: GORM AutoMigrate is DISABLED")
	}

	return db
}
