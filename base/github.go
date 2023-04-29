package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/models"
)

type GitHubRepo struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	HtmlUrl  string `json:"html_url"`
}

func getGitHubRepos(path string) ([]GitHubRepo, error) {
	var repos []GitHubRepo
	fullPath := GITHUB_API_BASE + path
	req, err := http.NewRequest("GET", fullPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+GITHUB_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(string(body))
		return nil, err
	}
	json.Unmarshal(body, &repos)

	return repos, nil
}

func concatCopyPreAllocate(slices [][]GitHubRepo) []GitHubRepo {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]GitHubRepo, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func getGitHubReposAll(paths []string) ([]GitHubRepo, error) {
	var allRepos [][]GitHubRepo
	for _, path := range paths {
		repos, err := getGitHubRepos(path)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		allRepos = append(allRepos, repos)
	}
	return concatCopyPreAllocate(allRepos), nil
}

func saveGitHubRepos(serviceId string, repos []GitHubRepo) error {
	bmiCollection, err := app.Dao().FindCollectionByNameOrId("bookmark_items")
	if err != nil {
		log.Fatal("Failed to get collection: bookmark_items")
		return err
	}
	for _, repo := range repos {
		record := models.NewRecord(bmiCollection)
		record.Set("name", repo.FullName)
		record.Set("url", repo.HtmlUrl)
		record.Set("service", serviceId)
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Fatal("Failed to save record")
			return err
		}
	}
	log.Println("Saved bookmarks for GitHub")
	return nil
}

type GitHubReposGetMultiplePayload struct {
	Paths []string `json:"paths"`
}

func saveBookmarksGitHubHandler(c echo.Context) error {
	record, err := app.Dao().FindFirstRecordByData("bookmark_services", "name", "GitHub")
	if err != nil {
		log.Println(err)
		return err
	}
	var payload GitHubReposGetMultiplePayload
	if err := c.Bind(&payload); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "bad request, failed to bind payload")
	}
	repos, err := getGitHubReposAll(payload.Paths)
	if err != nil {
		log.Fatal("Failed to get all GitHub repos")
		return err
	}
	if err = deleteBookmarksByService(record.Id); err != nil {
		log.Fatal("Failed to delete all GitHub repos")
		return err
	}
	if err = saveGitHubRepos(record.Id, repos); err != nil {
		log.Fatal("Failed to save all GitHub repos")
		return err
	}
	return c.String(http.StatusOK, "OK")
}
