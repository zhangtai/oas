package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/pocketbase/pocketbase/models"
)

func getTimestampWithTZ(e gocal.Event) (*time.Time, error) {
	var tzOffsite string
	startRawParsed := strings.Replace(e.Start.Format(time.RFC3339), "Z", "", 1)
	switch tz := e.RawStart.Params["TZID"]; tz {
	case "China Standard Time":
		tzOffsite = "+0800"
	case "India Standard Time":
		tzOffsite = "+0530"
	case "GMT Standard Time":
		tzOffsite = "+0100"
	default:
		tzOffsite = "+0000"
	}
	layout := "2006-01-02T15:04:05-0700"
	startInString := fmt.Sprintf("%s%s", startRawParsed, tzOffsite)
	eventStartWithTZ, err := time.Parse(layout, startInString)
	if err != nil {
		return nil, err
	}

	return &eventStartWithTZ, nil
}

type CalendarEntry struct {
	Event    gocal.Event
	Start    time.Time
	ZoomHost string
}

func getCalendarEntries(icsReader io.Reader, start time.Time, end time.Time) []CalendarEntry {
	c := gocal.NewParser(icsReader)
	c.Start, c.End = &start, &end
	c.Parse()
	var entries []CalendarEntry

	for _, e := range c.Events {
		var host string
		eventStartWithTZ, _ := getTimestampWithTZ(e)

		r, _ := regexp.Compile(`\\n\\n\\nHost:\\n\\n(\w.*?)\\n\\n`)
		match := r.FindStringSubmatch(e.Description)
		if len(match) > 0 {
			host = match[1]
		}

		if e.CustomAttributes["X-MICROSOFT-CDO-ALLDAYEVENT"] == "FALSE" && c.Start.Before(*eventStartWithTZ) {
			entries = append(entries, CalendarEntry{
				Event:    e,
				Start:    *eventStartWithTZ,
				ZoomHost: host,
			})
		}
	}
	return entries
}

func syncCalendar() {
	log.Println("Syncing calendar from ics")
	req, err := http.NewRequest("GET", CALENDAR_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	entries := getCalendarEntries(resp.Body, time.Now(), time.Now().Add(3*24*time.Hour))

	colName := "calendar_entries"
	q := app.DB().Builder.TruncateTable(colName)
	if _, err := q.Execute(); err != nil {
		log.Fatal(err)
	}

	col, err := app.Dao().FindCollectionByNameOrId(colName)
	if err != nil {
		log.Fatalf("Failed to get collection: %s", colName)
		log.Fatal(err)
	}
	for _, ce := range entries {
		record := models.NewRecord(col)
		record.Set("summary", ce.Event.Summary)
		record.Set("start", ce.Start)
		record.Set("zoomhost", ce.ZoomHost)
		record.Set("zoomurl", ce.Event.Location)
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Fatal("Failed to save record")
			log.Fatal(err)
		}
	}
}
