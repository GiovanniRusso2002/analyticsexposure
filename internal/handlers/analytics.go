package handlers

import (
	"net/http"

	"github.com/GiovanniRusso2002/analyticsexposure/internal/models"
	"github.com/GiovanniRusso2002/analyticsexposure/internal/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Handler holds the dependencies for HTTP handlers
type Handler struct {
	store storage.Store
}

// NewHandler creates a new handler
func NewHandler(store storage.Store) *Handler {
	return &Handler{
		store: store,
	}
}

// GetSubscriptions retrieves all subscriptions for an AF
func (h *Handler) GetSubscriptions(c echo.Context) error {
	afID := c.Param("afId")
	_ = c.QueryParam("supp-feat") // Optional parameter

	subs, err := h.store.GetAllSubscriptions(afID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, subs)
}

// CreateSubscription creates a new subscription
func (h *Handler) CreateSubscription(c echo.Context) error {
	afID := c.Param("afId")

	var req models.AnalyticsExposureSubsc
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Generate subscription ID
	subscriptionID := uuid.New().String()

	// Validate required fields
	if len(req.AnalyticsEventSubsc) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "analyEventsSubs is required"})
	}
	if req.NotifURI == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "notifUri is required"})
	}
	if req.NotifID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "notifId is required"})
	}

	// Create subscription
	if err := h.store.CreateSubscription(afID, subscriptionID, &req); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set("Location", "/"+afID+"/subscriptions/"+subscriptionID)
	return c.JSON(http.StatusCreated, req)
}

// GetSubscription retrieves a specific subscription
func (h *Handler) GetSubscription(c echo.Context) error {
	afID := c.Param("afId")
	subscriptionID := c.Param("subscriptionId")
	_ = c.QueryParam("supp-feat") // Optional parameter

	sub, err := h.store.GetSubscription(afID, subscriptionID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subscription not found"})
	}

	return c.JSON(http.StatusOK, sub)
}

// UpdateSubscription updates an existing subscription
func (h *Handler) UpdateSubscription(c echo.Context) error {
	afID := c.Param("afId")
	subscriptionID := c.Param("subscriptionId")

	var req models.AnalyticsExposureSubsc
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Validate required fields
	if len(req.AnalyticsEventSubsc) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "analyEventsSubs is required"})
	}
	if req.NotifURI == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "notifUri is required"})
	}
	if req.NotifID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "notifId is required"})
	}

	if err := h.store.UpdateSubscription(afID, subscriptionID, &req); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subscription not found"})
	}

	return c.JSON(http.StatusOK, req)
}

// DeleteSubscription deletes a subscription
func (h *Handler) DeleteSubscription(c echo.Context) error {
	afID := c.Param("afId")
	subscriptionID := c.Param("subscriptionId")

	if err := h.store.DeleteSubscription(afID, subscriptionID); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Subscription not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

// FetchAnalyticsData fetches analytics information
func (h *Handler) FetchAnalyticsData(c echo.Context) error {
	afID := c.Param("afId")

	var req models.AnalyticsRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Validate required fields
	if req.SuppFeat == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "suppFeat is required"})
	}

	data, err := h.store.GetAnalyticsData(afID, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if data == nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, data)
}
