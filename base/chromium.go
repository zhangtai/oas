package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

type BrowserTab struct {
	Title        string `json:"title"`
	Url          string `json:"url"`
	ScriptOutput string `json:"scriptOutput"`
}

type TabPayload struct {
	Script        string `json:"script"`
	UrlStartsWith string `json:"urlStartsWith"`
	Activate      bool   `json:"activate"`
}

func getChromiumTabByUrl(p TabPayload) (*BrowserTab, error) {
	jaxScript := fmt.Sprintf(
		`const findTabByUrlStarts = (urlStartsWith) => { for (const window of Application("Chromium").windows()) { for (const tab of window.tabs()) { if (tab.url().startsWith(urlStartsWith)) { return tab } } } return null; };
	const getTab = (tab, script) => { output = script ? tab.execute({ javascript: script }) : ""; return JSON.stringify({ title: tab.title(), url: tab.url(), scriptOutput: output.toString() }) };
	(function (){
		let urlStartsWith = "%s";
		let script = %s;
		tab = findTabByUrlStarts(urlStartsWith);
		return tab ? getTab(tab, script) : null;
	})()`,
		p.UrlStartsWith,
		fmt.Sprintf("`%s`", p.Script),
	)
	output, err := runAppleScript(AppleScriptPayload{
		Script: jaxScript,
		IsJavaScript: true,
	})
	// log.Println(output)
	if err != nil {
		return nil, err
	}
	var tab BrowserTab
	err = json.Unmarshal([]byte(output), &tab)
	if err != nil {
		return nil, err
	}
	return &tab, nil
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
