package main

import (
	"log"
	"net/http"

	"github.com/GiovanniRusso2002/analyticsexposure/internal/handlers"
	"github.com/GiovanniRusso2002/analyticsexposure/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultPort = "8080"
	apiVersion  = "v1"
	basePath    = "/3gpp-analyticsexposure/" + apiVersion
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Configure middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize storage
	store := storage.NewInMemoryStore()

	// Initialize handlers
	h := handlers.NewHandler(store)

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "UP",
			"service": "3GPP Analytics Exposure API",
			"version": apiVersion,
		})
	})

	// OpenAPI documentation endpoint
	e.GET("/openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]string{
				"title":   "3GPP Analytics Exposure API",
				"version": apiVersion,
			},
			"servers": []map[string]string{
				{
					"url": basePath,
				},
			},
		})
	})

	// Analytics Exposure Subscriptions
	// GET /{afId}/subscriptions
	e.GET(basePath+"/:afId/subscriptions", h.GetSubscriptions)

	// POST /{afId}/subscriptions
	e.POST(basePath+"/:afId/subscriptions", h.CreateSubscription)

	// GET /{afId}/subscriptions/{subscriptionId}
	e.GET(basePath+"/:afId/subscriptions/:subscriptionId", h.GetSubscription)

	// PUT /{afId}/subscriptions/{subscriptionId}
	e.PUT(basePath+"/:afId/subscriptions/:subscriptionId", h.UpdateSubscription)

	// DELETE /{afId}/subscriptions/{subscriptionId}
	e.DELETE(basePath+"/:afId/subscriptions/:subscriptionId", h.DeleteSubscription)

	// POST /{afId}/fetch
	e.POST(basePath+"/:afId/fetch", h.FetchAnalyticsData)

	// Start server
	port := defaultPort
	address := ":" + port

	log.Printf("Starting 3GPP Analytics Exposure API server on http://localhost%s%s\n", address, basePath)
	log.Printf("Health check available at http://localhost%s/health\n", address)
	log.Printf("OpenAPI spec available at http://localhost%s/openapi.json\n", address)

	if err := e.Start(address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
