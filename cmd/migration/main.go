package main

import (
	"database/sql"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database/postgres"
	"github.com/samuelralmeida/product-catalog-api/migrations"

	"github.com/pressly/goose/v3"
)

func main() {
	conn, err := postgres.Open(postgres.DefaultConfig())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = Migrate(conn)
	if err != nil {
		panic(err)
	}

}

func migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func Migrate(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	// undo filesystem to prevent errors
	defer func() { goose.SetBaseFS(nil) }()
	return migrate(db, ".")
}
