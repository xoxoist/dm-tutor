package main

import (
	"flag"
	"github.com/xoxoist/dm-tutor/database"
)

func main() {
	action := flag.String("action", "UP", "desired action for your database migration")
	flag.Parse()

	mg, err := database.NewMigration(database.MySQLBuilder, database.Config{
		DatabaseHost: "localhost",
		DatabasePort: "3306",
		DatabaseName: "md_tutor",
		DatabaseUser: "root",
		DatabasePasw: "root",
		DatabaseDrvr: "mysql",
		DatabaseMdir: "database/migration",
	})
	if err != nil {
		panic(err)
	}

	err = mg.Action(*action)
	if err != nil {
		panic(err)
	}
}
