package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/robfig/cron/v3"
)

var (
	app             *pocketbase.PocketBase = pocketbase.New()
	CONFLUENCE_BASE string
	CONFLUENCE_USER string
	CONFLUENCE_PASS string
	GITHUB_API_BASE string
	GITHUB_TOKEN    string
	REMINDERS_CLI   string
	CALENDAR_URL    string
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to get home dir for the user")
	}
	err = godotenv.Load(fmt.Sprintf("%s/.local/var/oas/.env", homeDir))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CONFLUENCE_BASE = os.Getenv("CONFLUENCE_BASE")
	CONFLUENCE_USER = os.Getenv("CONFLUENCE_USER")
	CONFLUENCE_PASS = os.Getenv("CONFLUENCE_PASS")
	GITHUB_API_BASE = os.Getenv("GITHUB_API_BASE")
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	REMINDERS_CLI = os.Getenv("REMINDERS_CLI")
	CALENDAR_URL = os.Getenv("CALENDAR_URL")

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		addRoutes(e)
		return nil
	})

	c := cron.New()
	c.AddFunc("@every 15m", syncCalendar)
	c.Start()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
