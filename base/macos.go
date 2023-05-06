package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v5"
)

type AppleScriptPayload struct {
	Script       string `json:"script"`
	IsJavaScript bool   `json:"isJavaScript"`
}

type CommandPayload struct {
	Command []string `json:"command"`
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

func runCommand(p CommandPayload) (string, error) {
	cmd := exec.Command(p.Command[0], p.Command[1:]...)
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

func commandHandler(c echo.Context) error {
	var payload CommandPayload
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	output, err := runCommand(payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	return c.String(http.StatusOK, output)
}
