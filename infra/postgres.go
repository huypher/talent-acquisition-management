package infra

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgres(cfg *Config) (*bun.DB, func(), error) {

	dsn := "postgres://dobi:dobi@host.docker.internal:5432/dobi-oms?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return db, func() {
		db.Close()
	}, nil
}
