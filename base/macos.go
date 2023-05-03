package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v5"
)

type OpenPayload struct {
	Uri string `json:"uri"`
}

type AppleScriptPayload struct {
	Script string `json:"script"`
	IsJavaScript bool `json:"isJavaScript"`
}

func runAppleScript(p AppleScriptPayload) (string, error) {
	cmd := exec.Command("/usr/bin/osascript", "-s", "h", "-e", p.Script)
	if p.IsJavaScript {
		cmd = exec.Command("/usr/bin/osascript", "-s", "h", "-l", "JavaScript", "-e", p.Script)
	}

	var out strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func open(uri string) (string, error) {
	cmd := exec.Command("/usr/bin/open", uri)
	var out strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func systemOpenHandler(c echo.Context) error {
	var payload OpenPayload
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	output, err := open(payload.Uri)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}

	return c.String(http.StatusOK, output)
}

func appleScriptHandler(c echo.Context) error {
	var payload AppleScriptPayload
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	output, err := runAppleScript(payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	return c.String(http.StatusOK, output)
}
