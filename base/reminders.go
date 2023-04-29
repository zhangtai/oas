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
	CalendarTitle    string        `json:"calendarTitle"`
	Id               string        `json:"id"`
	Title            string        `json:"title"`
	CreationDate     time.Time     `json:"creationDate"`
	StartDate        *time.Time    `json:"startDate"`
	CompletionDate   *time.Time    `json:"completionDate"`
	LastModifiedDate time.Time     `json:"lastModifiedDate"`
	DueDate          *time.Time    `json:"dueDate"`
	Notes            *string       `json:"notes"`
	Meta             *ReminderMeta `json:"meta"`
	Priority         int           `json:"priority"`
	IsCompleted      bool          `json:"isCompleted"`
}

type RemindersAll map[string][]Reminder

type ListNotFoundError struct {
	Name string
}

func (e *ListNotFoundError) Error() string {
	return fmt.Sprintf("list not found: %v", e.Name)
}

func parseReminderNotes(input string) (*ReminderMeta, error) {
	var meta ReminderMeta
	if err := json.Unmarshal([]byte(input), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func getRemindersByList(listName string) ([]Reminder, error) {
	cmd := exec.Command("/Users/taizhang/.local/bin/reminders", "export-all")
	var out strings.Builder
	var stderr strings.Builder
	var rl RemindersAll
	var reminders []Reminder
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return rl[listName], err
	}
	if err = json.Unmarshal([]byte(out.String()), &rl); err != nil {
		return nil, err
	}
	remindersResp, ok := rl[listName]
	if ok {
		for _, r := range remindersResp {
			if r.Notes != nil {
				meta, err := parseReminderNotes(*r.Notes)
				if err != nil {
					log.Println("failed to parse notes to meta for reminder", r.Title)
				}
				r.Meta = meta
			}
			reminders = append(reminders, r)
		}
	} else {
		return nil, &ListNotFoundError{Name: listName}
	}
	return reminders, nil
}

func remindersListGetHandler(c echo.Context) error {
	listName := c.PathParam("name")
	reminders, err := getRemindersByList(listName)
	if err != nil {
		switch err.(type) {
		case *ListNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			log.Println(err)
			return c.String(http.StatusInternalServerError, "NotOK")
		}
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
	if getNext == true {
		if len(reminders) > 1 {
			return c.JSON(http.StatusOK, reminders[0])
		}
	}

	return c.JSON(http.StatusOK, reminders)
}
