package main

import (
	"log"

	"github.com/pocketbase/dbx"
)

func deleteBookmarksByService(serviceId string) error {
	records, err := app.Dao().FindRecordsByExpr("bookmark_items", dbx.HashExp{"service": serviceId})
	if err != nil {
		log.Fatal("Failed to get records of bookmark_items of service")
		return err
	}

	for _, record := range records {
		if err := app.Dao().DeleteRecord(record); err != nil {
			log.Fatal("Failed to delete record")
			log.Fatal(record)
			return err
		}
	}
	return nil
}
