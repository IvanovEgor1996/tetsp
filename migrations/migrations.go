package migrations

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

var connConf = ConnectConf{
	Host:   os.Getenv("PG_HOST"),
	Port:   os.Getenv("PG_PORT"),
	User:   os.Getenv("PG_USERNAME"),
	Pass:   os.Getenv("PG_PASSWORD"),
	DBName: os.Getenv("PG_DBNAME"),
}

type ConnectConf struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"db_name"`
}

func Migrate() {

	flag.Usage = usage
	flag.Parse()

	db := pg.Connect(&pg.Options{
		Addr:     connConf.Host + ":" + connConf.Port,
		User:     connConf.User,
		Password: connConf.Pass,
		Database: connConf.DBName,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
