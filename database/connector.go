package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// connect database connector using builtin go sql library
func connect(dialect, connStr string) (*sql.DB, error) {
	db, err := sql.Open(dialect, connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MySQLBuilder decorator function for constructing mysql connection string
func MySQLBuilder(cfg Config) (*migrate.Migrate, error) {
	// construct connection string and connect
	const formatString = "%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true"
	credentials := []any{
		cfg.DatabaseUser, cfg.DatabasePasw, cfg.DatabaseHost,
		cfg.DatabasePort, cfg.DatabaseName,
	}
	finalCsf := fmt.Sprintf(formatString, credentials...)
	db, err := connect("mysql", finalCsf)
	if err != nil {
		return nil, err
	}

	// create migrate instance from connected database client
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("file://%s", cfg.DatabaseMdir)
	mgrt, err := migrate.NewWithDatabaseInstance(filePath, cfg.DatabaseDrvr, driver)
	if err != nil {
		return nil, err
	}
	return mgrt, nil
}

// PostgresBuilder decorator function for constructing mysql connection string
func PostgresBuilder(cfg Config) (*migrate.Migrate, error) {
	// construct connection string and connect
	const formatString = "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	credentials := []any{
		cfg.DatabaseUser, cfg.DatabasePasw, cfg.DatabaseHost,
		cfg.DatabasePort, cfg.DatabaseName,
	}
	finalCsf := fmt.Sprintf(formatString, credentials...)
	db, err := connect("postgres", finalCsf)
	if err != nil {
		return nil, err
	}

	// create migrate instance from connected database client
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("file://%s", cfg.DatabaseMdir)
	mgrt, err := migrate.NewWithDatabaseInstance(filePath, cfg.DatabaseDrvr, driver)
	if err != nil {
		return nil, err
	}
	return mgrt, nil
}
