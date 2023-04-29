package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var (
	app             *pocketbase.PocketBase = pocketbase.New()
	AD_USER         string
	AD_PASS         string
	GITHUB_API_BASE string
	GITHUB_TOKEN    string
	REMINDERS_CLI   string
)

func main() {
	err := godotenv.Load("/Users/taizhang/.local/var/oas/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AD_USER = os.Getenv("AD_USER")
	AD_PASS = os.Getenv("AD_PASS")
	GITHUB_API_BASE = os.Getenv("GITHUB_API_BASE")
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	REMINDERS_CLI = os.Getenv("REMINDERS_CLI")

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		addRoutes(e)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
