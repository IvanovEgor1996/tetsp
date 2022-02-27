package main

import (
	"os"
	"testP/migrations"
	"testP/service"
)

func main() {

	switch os.Getenv("APP_MODE") {
	case "MIGRATE":
		migrations.Migrate()
	default:
		serv, err := service.NewService()
		if err != nil {
			panic(err)
		}
		serv.Run()
	}

}
