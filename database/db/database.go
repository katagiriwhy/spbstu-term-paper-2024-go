package db

import (
	"fmt"

	"github.com/katagiriwhy/database/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dbCfg := cfg.DB
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	postDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = postDB.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
