package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
)

type ReminderMeta struct {
	Host string `json:"host"`
	Url  string `json:"url"`
}

type Reminder struct {
	DueDate     *time.Time    `json:"dueDate"`
	ExternalId  string        `json:"externalId"`
	IsCompleted bool          `json:"isCompleted"`
	List        string        `json:"list"`
	Notes       *string       `json:"notes"`
	Priority    int           `json:"priority"`
	Title       string        `json:"title"`
	Meta        *ReminderMeta `json:"meta"`
}

type RemindersShowAllFilter struct {
	DueDate time.Time `json:"dueDate"`
}

func getRemindersAll(filter RemindersShowAllFilter) ([]Reminder, error) {
	var cmd *exec.Cmd
	if !filter.DueDate.IsZero() {
		log.Printf("Getting reminders with due date: %s", filter.DueDate.Format("2006-01-02"))
		cmd = exec.Command(REMINDERS_CLI, "show-all", "-f", "json", "-d", filter.DueDate.Format("2006-01-02"))
	} else {
		cmd = exec.Command(REMINDERS_CLI, "show-all", "-f", "json")
	}
	var out strings.Builder
	var stderr strings.Builder
	var remindersResponse []Reminder
	var reminders []Reminder
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	log.Println(out.String())
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return nil, err
	}
	if err = json.Unmarshal([]byte(out.String()), &remindersResponse); err != nil {
		return nil, err
	}
	for _, r := range remindersResponse {
		if r.Notes != nil {
			meta, err := parseReminderNotes(*r.Notes)
			if err != nil {
				log.Println("failed to parse notes to meta for reminder", r.Title)
			}
			r.Meta = meta
		}
		reminders = append(reminders, r)
	}
	return reminders, nil
}

func parseReminderNotes(input string) (*ReminderMeta, error) {
	var meta ReminderMeta
	if err := json.Unmarshal([]byte(input), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func getRemindersAllHandler(c echo.Context) error {
	var filter RemindersShowAllFilter
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	reminders, err := getRemindersAll(filter)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "NotOK")
	}
	sort.Slice(reminders, func(i, j int) bool {
		if reminders[i].DueDate == nil {
			return false
		}
		if reminders[j].DueDate == nil {
			return true
		}
		return reminders[i].DueDate.Unix() < reminders[j].DueDate.Unix()
	})

	var getNext bool
	getNextStr := c.QueryParam("next")
	getNext, err = strconv.ParseBool(getNextStr)
	if err != nil {
		getNext = false
	}
	if getNext {
		if len(reminders) > 1 {
			return c.JSON(http.StatusOK, reminders[0])
		}
	}

	return c.JSON(http.StatusOK, reminders)
}
