package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ez-connect/go-rest/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	appName    = "Example"
	appVersion = "v0.0.1"
)

func main() {
	/// Server
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status}: ${method} ${latency_human} ${uri} ${remote_ip} ${error}\n",
	}))
	e.Use(middleware.Recover())

	// CORS default
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			// http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		ExposeHeaders: []string{
			rest.HeaderTotalCount,
		},
	}))

	// Route => handler
	// routes.Init(e, driver.Get())

	// Start server
	addr := fmt.Sprintf("%s:%v", "", 4000)
	fmt.Println(appName, appVersion, "serve at", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
