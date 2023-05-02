package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
)

type BrowserTab struct {
	Index        int    `json:"index"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ScriptOutput string `json:"scriptOutput"`
}

func getChromiumTabByIndex(index int) (BrowserTab, error) {
	script := fmt.Sprintf(`tell application "Chromium"
		set theTab to tab %d of window 1
		return %d & title of theTab & URL of theTab
	end tell`, index, index)
	output, err := runAppleScript(script)
	tabComponents := strings.Split(output, ", ")
	if err != nil || len(tabComponents) != 3 {
		log.Println(err)
		return BrowserTab{
			Index: -1,
			Title: "",
			Url:   "",
		}, err
	}
	return BrowserTab{
		Index: index,
		Title: tabComponents[1],
		Url:   tabComponents[2],
	}, nil
}

func getChromiumTabByUrl(p TabPayload) (BrowserTab, error) {
	script := fmt.Sprintf(`tell application "Chromium"
	set executeJs to %t
	set activateFoundTab to %t
	repeat with the_window in every window
		set tab_index to 0
		repeat with the_tab in every tab in the_window
		    set tab_index to tab_index + 1
			set the_url to the URL of the_tab
			set the_title to the title of the_tab
			set jsOutput to "None"
			if {the_url starts with "%s"} then
				if activateFoundTab then
					set active tab index of the_window to tab_index
					activate the_window
				end if
				if executeJs then
					set jsOutput to execute tab tab_index of the_window javascript "%s"
				end if
		        return tab_index & the_title & the_url & jsOutput
			end if
		end repeat
	end repeat
	end tell`,
		p.Script != "",
		p.Activate,
		p.UrlStartsWith,
		p.Script,
	)
	output, err := runAppleScript(script)
	log.Println(output)
	tabComponents := strings.Split(output, ", ")
	if err != nil || len(tabComponents) != 4 {
		log.Println(err)
		return BrowserTab{
			Index:        -1,
			Title:        "",
			Url:          "",
			ScriptOutput: "",
		}, err
	}
	index, err := strconv.Atoi(tabComponents[0])
	if err != nil {
		log.Println(err)
		return BrowserTab{
			Index:        -1,
			Title:        "",
			Url:          "",
			ScriptOutput: "",
		}, err
	}
	return BrowserTab{
		Index:        index,
		Title:        tabComponents[1],
		Url:          tabComponents[2],
		ScriptOutput: tabComponents[3],
	}, nil
}

func executeJavaScript(tab int, js string) (string, error) {
	script := fmt.Sprintf(`tell application "Chromium"
	    execute tab %d of first window javascript "%s"
	end tell`, tab, js)
	output, err := runAppleScript(script)
	if err != nil {
		return "", err
	}
	return output, nil
}

func activateChromiumTab(tab int) error {
	script := fmt.Sprintf(`tell application "Chromium"
		if not running then
			run
			delay 0.25
		end if
		set active tab index of first window to %d
		activate
	end tell`, tab)
	if _, err := runAppleScript(script); err != nil {
		return err
	}
	log.Printf("Activated tab %d of Chromium", tab)
	return nil
}

func chromiumTabsHandler(c echo.Context) error {
	index, err := strconv.Atoi(c.PathParam("index"))
	if err != nil {
		return c.String(http.StatusBadRequest, "not a valid tab index")
	}
	tab, err := getChromiumTabByIndex(index)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to get tab by index")
	}

	return c.JSON(http.StatusOK, tab)
}

type TabPayload struct {
	Script        string `json:"script"`
	UrlStartsWith string `json:"urlStartsWith"`
	Activate      bool   `json:"activate"`
}

func chromiumTabsActionsHandler(c echo.Context) error {
	index, err := strconv.Atoi(c.PathParam("index"))
	if err != nil {
		return c.String(http.StatusBadRequest, "not a valid tab index")
	}
	action := c.PathParam("action")

	if action == "activate" {
		if err := activateChromiumTab(index); err != nil {
			return c.String(http.StatusServiceUnavailable, "NotOK")
		}
	}

	if action == "execute" {
		var tp TabPayload
		if err := c.Bind(&tp); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
		}
		output, err := executeJavaScript(index, tp.Script)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "bad request, failed to execute javascript")
		}
		return c.String(http.StatusOK, output)
	}

	return c.String(http.StatusOK, "OK")
}

func chromiumTabFindHandler(c echo.Context) error {
	var p TabPayload
	if err := c.Bind(&p); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	tab, err := getChromiumTabByUrl(p)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "bad request, failed to execute javascript")
	}
	return c.JSON(http.StatusOK, tab)
}