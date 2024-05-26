package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

type (
	// ConnectionFunc decorator function for constructing connection string
	ConnectionFunc func(cfg Config) (*migrate.Migrate, error)

	// Config database configuration detail for connecting to desired database server
	Config struct {
		DatabaseHost string // database host
		DatabasePort string // database port
		DatabaseName string // database name
		DatabaseUser string // database user
		DatabasePasw string // database password
		DatabaseDrvr string // database driver
		DatabaseMdir string // database migration dir
	}

	// Migration interface for migration
	Migration interface {
		Action(name string) error
		down() error
		up() error
	}

	// migrationImpl implementation struct
	migrationImpl struct {
		m *migrate.Migrate
	}
)

func NewMigration(cf ConnectionFunc, cfg Config) (Migration, error) {
	// initialize migration instance
	m, err := cf(cfg)
	if err != nil {
		return nil, err
	}
	m.Log = NewLog(
		log.New(log.Writer(),
			"migration: ",
			log.LstdFlags|log.Lshortfile,
		),
	)
	mg := &migrationImpl{m: m}
	return mg, nil
}

// Action migrating desired migration version start constructing database
func (mgr *migrationImpl) Action(name string) error {
	mgr.m.Log.Printf("migaration.Action: %s", "starting database migration")
	switch name {
	case "UP":
		return mgr.up()
	case "DOWN":
		return mgr.down()
	default:
		mgr.m.Log.Printf("migaration.Action: %s (%s)", "unknown action name", name)
		return errors.New("unknown action")
	}
}

// down performing all SQL inside `down` migration files
func (mgr *migrationImpl) down() error {
	err := mgr.m.Down()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			mgr.m.Log.Printf("migaration.down: %s (%s)", "deconstructing database done", err.Error())
			return nil
		}
		mgr.m.Log.Printf("migaration.down: %s (%s)", "failed to deconstructing database", err.Error())
		return err
	}
	return nil
}

// up performing all SQL inside `up` migration files
func (mgr *migrationImpl) up() error {
	mgr.m.Log.Verbose()
	err := mgr.m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			mgr.m.Log.Printf("migaration.down: %s (%s)", "constructing database done", err.Error())
			return nil
		}
		mgr.m.Log.Printf("migaration.down: %s (%s)", "failed to constructing database", err.Error())
		return err
	}
	return nil
}
