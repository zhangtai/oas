package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/models"
)

type ConfluencePageLinks struct {
	Tinyui string `json:"tinyui"`
	Base   string `json:"base"`
}

type ConfluencePage struct {
	Id     string              `json:"id"`
	Title  string              `json:"title"`
	Type   string              `json:"type"`
	Status string              `json:"status"`
	Links  ConfluencePageLinks `json:"_links"`
}

type ConfluenceResponse struct {
	Results []ConfluencePage `json:"results"`
}

type ConfluenceQueryResponse struct {
	Results []ConfluencePage    `json:"results"`
	Links   ConfluencePageLinks `json:"_links"`
}

type PageAncestor struct {
	Id int `json:"id"`
}

type ConfluenceSpace struct {
	Key string `json:"key"`
}

type PageStorage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type PageStorageHolder struct {
	Storage PageStorage `json:"storage"`
}

type ConfluencePagePayload struct {
	Type      string            `json:"type"`
	Title     string            `json:"title"`
	Ancestors []PageAncestor    `json:"ancestors"`
	Space     ConfluenceSpace   `json:"space"`
	Body      PageStorageHolder `json:"body"`
}

type SimplePagePayload struct {
	SpaceKey string `json:"spaceKey"`
	Parent   int    `json:"parent"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type ConfluencePageCreatedResponse struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

func getConfluenceFavs() (*ConfluenceQueryResponse, error) {
	var confluence ConfluenceQueryResponse
	endpoint := CONFLUENCE_BASE + "/rest/api/content/search?limit=50&cql=favourite=currentUser()"
	req, err := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(CONFLUENCE_USER, CONFLUENCE_PASS)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &confluence)
	log.Printf("Got confluence faviroute pages: %d", len(confluence.Results))
	return &confluence, nil
}

func saveConfluenceItems(serviceId string, c ConfluenceQueryResponse) error {
	bmiCollection, err := app.Dao().FindCollectionByNameOrId("bookmark_items")
	if err != nil {
		log.Fatal("Failed to get collection: bookmark_items")
		return err
	}
	for _, item := range c.Results {
		record := models.NewRecord(bmiCollection)
		record.Set("name", item.Title)
		record.Set("url", c.Links.Base+item.Links.Tinyui)
		record.Set("service", serviceId)
		if err := app.Dao().SaveRecord(record); err != nil {
			log.Fatal("Failed to save record")
			return err
		}
	}
	log.Println("Saved bookmarks for Confluence")
	return nil
}

func saveBookmarksConfluence(c echo.Context) error {
	record, err := app.Dao().FindFirstRecordByData("bookmark_services", "name", "Confluence")
	if err != nil {
		log.Fatal("Failed to get Confluence record")
		return err
	}
	confluenceResponse, err := getConfluenceFavs()
	if err != nil {
		log.Fatal("Failed to get all Confluence items")
		return err
	}
	if err = deleteBookmarksByService(record.Id); err != nil {
		log.Fatal("Failed to delete all Confluence items")
		return err
	}
	if err = saveConfluenceItems(record.Id, *confluenceResponse); err != nil {
		log.Fatal("Failed to save all Confluence items")
		return err
	}
	return c.String(http.StatusOK, "OK")
}

func createConfluencePagesHandler(c echo.Context) error {
	var spps []SimplePagePayload
	var results []string
	if err := c.Bind(&spps); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	for _, spp := range spps {
		result, err := createConfluencePage(spp)
		if err != nil {
			return err
		}
		results = append(results, result.Links.Base+result.Links.Tinyui)
	}
	return c.JSON(http.StatusOK, results)
}

func createConfluencePage(p SimplePagePayload) (*ConfluencePage, error) {
	endpoint := CONFLUENCE_BASE + "/rest/api/content/"
	payload := &ConfluencePagePayload{
		Type:  "page",
		Title: p.Title,
		Ancestors: []PageAncestor{{
			Id: p.Parent,
		}},
		Space: ConfluenceSpace{
			Key: p.SpaceKey,
		},
		Body: PageStorageHolder{
			Storage: PageStorage{
				Value:          p.Content,
				Representation: "storage",
			},
		},
	}
	log.Println(payload)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)
	req, err := http.NewRequest("POST", endpoint, payloadBuf)
	req.Header.Set("X-Atlassian-Token", "no-check")
	req.Header.Set("content-type", "application/json")
	req.SetBasicAuth(CONFLUENCE_USER, CONFLUENCE_PASS)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Println("response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Println(string(body))
		return nil, errors.New("not created")
	}
	var created ConfluencePage
	json.Unmarshal(body, &created)
	return &created, nil
}
