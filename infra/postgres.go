package infra

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	migration "github.com/pghuy/talent-acquisition-management/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *Config) (*gorm.DB, func(), error) {
	databaseUrl := "postgres://tam:tam@postgres:5432/talent-acquisition-management?sslmode=disable"
	sqldb := postgres.Open(databaseUrl)
	db, err := gorm.Open(sqldb, &gorm.Config{})

	postgresInstance, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	if err := postgresInstance.Ping(); err != nil {
		postgresInstance.Close()
		return nil, nil, err
	}

	err = migration.MigrateUp(postgresInstance)
	if err != nil {
		postgresInstance.Close()
		return nil, nil, err
	}

	return db, func() {
		postgresInstance.Close()
	}, nil
}
