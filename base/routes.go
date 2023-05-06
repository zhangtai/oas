package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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

func addRoutes(e *core.ServeEvent) {
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/system/command",
		Handler: commandHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/system/applescript",
		Handler: appleScriptHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/apps/chromium/tab",
		Handler: chromiumTabFindHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/api/apps/reminders/lists/:name",
		Handler: remindersListGetHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/services/github/bookmarks",
		Handler: saveBookmarksGitHubHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/services/confluence/bookmarks",
		Handler: saveBookmarksConfluence,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
	e.Router.AddRoute(echo.Route{
		Method:  http.MethodPost,
		Path:    "/api/services/confluence/pages",
		Handler: createConfluencePagesHandler,
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(app),
		},
	})
}
