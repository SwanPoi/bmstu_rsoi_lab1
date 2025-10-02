package db

import (
	"fmt"
	"net/url"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"

	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"github.com/SwanPoi/bmstu_rsoi_lab1/pkg/logger"
	"github.com/SwanPoi/bmstu_rsoi_lab1/config"
)

func GetConnectionString(cfg *config.DatabaseConfig) (connStr string) {
	 dsn := url.URL{
        Scheme:   cfg.Driver,
        User:     url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
        Host:     fmt.Sprintf("%s:%d", cfg.Postgres.Host, cfg.Postgres.Port),
        Path:     cfg.Postgres.Database,
    }

	q := dsn.Query()
	q.Set("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	return dsn.String()
	// return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
	// 	cfg.Driver,
	// 	cfg.Postgres.User,
	// 	cfg.Postgres.Password,
	// 	cfg.Postgres.Host,
	// 	cfg.Postgres.Port,
	// 	cfg.Postgres.Database,
	// )
}

func Init(url string, logger *logger.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logger.Fatalf(err.Error())
	}

	db.AutoMigrate(&models.Person{})

	return db
}