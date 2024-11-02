// handler/event_handler.go
package handler

import (
	"api-test/models"
	"api-test/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService *services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Date        string `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	event := models.Event{
		Name:        input.Name,
		Description: input.Description,
		Date:        date,
		CreatedBy:   c.GetUint("user_id"), 
	}

	if err := h.eventService.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully"})
}
