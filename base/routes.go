package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

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
		Method:  http.MethodPost,
		Path:    "/api/apps/reminders/all",
		Handler: getRemindersAllHandler,
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
