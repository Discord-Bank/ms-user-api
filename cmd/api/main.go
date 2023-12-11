package main

import (
	"log"
	"ms-user-api/config"
	"ms-user-api/internal/http/router"
	"ms-user-api/user/db"
	"os"
	"time"
)

const TIMEOUT = 30 * time.Second

func main() {
	envs, err := config.LoadEnvVars()
	if err != nil {
		log.Fatalln("Failed loading env", err)
		os.Exit(1)
	}
	orm, err := db.NewDatabase(envs.DSN)
	if err != nil {
		log.Fatal("internal server error, " + err.Error())
		os.Exit(1)
	}
	orm.AutoMigrateSetup()
	h := router.Handlers(orm)
	err = h.Start(envs.APIPort)

	if err != nil {
		log.Fatal("Error starting api, error: " + err.Error())
		os.Exit(1)
	}
}
