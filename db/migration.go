package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func MigrateUp(instance *sql.DB) error {
	driver, err := postgres.WithInstance(instance, &postgres.Config{})
	if err != nil {
		log.Printf("get db instance err=%v\n", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Printf("get db instance err=%v\n", err)
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("migration no changes")
			return nil
		}

		log.Printf("migrate up err=%v\n", err)
		return err
	}

	log.Println("migration success")

	return nil
}
